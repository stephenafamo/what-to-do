package controller

import (
	"net/http"
	// "fmt"
	. "github.com/stephenafamo/what-to-do/models"
)

type IndexController struct {
	Controller
}

func (i *IndexController) Index(w http.ResponseWriter, r *http.Request, p interface{}) {
	completed := []Todo{}
	todo := []Todo{}
	DB.Where("completed_on is not null").Order("completed_on desc").Limit(5).Find(&completed)
	DB.Where("completed_on is null").Find(&todo)
	data:= struct {
        Status string
        Completed []Todo
        Todo []Todo
        Vars map[string]string
    }{
    	Status: "success",
    	Completed: completed,
    	Todo: todo,
    	Vars: i.GetVars(r),
    }
	w.WriteHeader(http.StatusOK)
	i.Render(w, "index", data)
}
