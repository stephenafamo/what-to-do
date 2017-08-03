package controller

import (
	"net/http"
	// "time"
	"encoding/json"
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
