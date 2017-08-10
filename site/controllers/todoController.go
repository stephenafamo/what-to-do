package controller

import (
	"net/http"
	"fmt"
	"time"
	"strconv"
	// "io/ioutil"
	"encoding/json"
	"database/sql"
	"github.com/jinzhu/gorm"
	. "github.com/stephenafamo/what-to-do/models"
)

type TodoController struct {
	Controller
}

func (i *TodoController) Add(w http.ResponseWriter, r *http.Request, p interface {}) {
	// r.ParseMultipartForm(1024)
	r.ParseForm()
	title := r.FormValue("title")
	if title == "" {
		return	
	}
	body, _ := time.Parse("2006-01-02 15:04", "2017-01-02 11:11")
	// body = time.Now()
	fmt.Println(body)
	fmt.Println(time.Now())
	todo := Todo{Title: title}
	DB.Create(&todo)
	jData, _ := json.Marshal(struct {
        Status string
        Data Todo
    }{
    	Status: "success",
    	Data: todo,
    })
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (i *TodoController) Edit(w http.ResponseWriter, r *http.Request, p interface {}) {
	r.ParseForm()
	todo := Todo{}
	vars := i.GetVars(r)
	DB.First(&todo, vars["ID"])

	oldTime, _ := time.Parse("2006-01-02 15:04:05", "1999-01-02 15:04:05")
	newValues := make(map[string]interface{})
	for key, value := range r.Form {

		switch key {
			case "title":
				newValues[key] = value[0]
			case "description":
				newValues[key] = sql.NullString{String: value[0], Valid: true}
			case "priority":
				newValues[key], _ = strconv.Atoi(value[0])
			case "estimated_duration":
				estimated_duration, _ := strconv.ParseInt(value[0], 10, 64)
				newValues[key] = sql.NullInt64{Int64: estimated_duration, Valid: true}
			case "actual_duration":
				actual_duration, _ := strconv.ParseInt(value[0], 10, 64)
				newValues[key] = sql.NullInt64{Int64: actual_duration, Valid: true}
			case "do_on":
				theTime, _ := time.Parse("2006-01-02 15:04:05", value[0])
				if theTime.After(oldTime) {
					newValues[key] = theTime
				}
			case "do_before":
				theTime, _ := time.Parse("2006-01-02 15:04:05", value[0])
				if theTime.After(oldTime) {
					newValues[key] = theTime
				}
			case "completed_on":
				theTime, _ := time.Parse("2006-01-02 15:04:05", value[0])
				if theTime.After(oldTime) {
					newValues[key] = theTime
				}
				if value[0] == "now" {
					newValues[key] = gorm.NowFunc()
				}
		}

		if value[0] == "null" {
			newValues[key] = nil			
		}
	}
	DB.Model(&todo).Updates(newValues)

	fmt.Println(r.Form)
	fmt.Println(newValues)
	fmt.Println(todo)
	jData, _ := json.Marshal(struct {
        Status string
        Data Todo
    }{
    	Status: "success",
    	Data: todo,
    })
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (i *TodoController) Complete(w http.ResponseWriter, r *http.Request, p interface {}) {
	// r.ParseMultipartForm(1024)
	todo := Todo{}
	vars := i.GetVars(r)
	DB.First(&todo, vars["ID"])
	DB.Model(&todo).Update("completed_on", gorm.NowFunc())
	jData, _ := json.Marshal(struct {
        Status string
        Data Todo
    }{
    	Status: "success",
    	Data: todo,
    })
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (i *TodoController) Delete(w http.ResponseWriter, r *http.Request, p interface {}) {
	// r.ParseMultipartForm(1024)
	todo := Todo{}
	vars := i.GetVars(r)
	DB.First(&todo, vars["ID"])
	DB.Delete(&todo)
	jData, _ := json.Marshal(struct {
        Status string
        Data Todo
    }{
    	Status: "success",
    	Data: todo,
    })
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}
