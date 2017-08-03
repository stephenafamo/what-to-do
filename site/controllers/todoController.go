package controller

import (
	"net/http"
	"encoding/json"
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
