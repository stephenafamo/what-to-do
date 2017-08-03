package controller

import (
	"net/http"
	"fmt"
	. "github.com/stephenafamo/what-to-do/models"
)

type IndexController struct {
	Controller
}

func (i *IndexController) Index(w http.ResponseWriter, r *http.Request, p interface{}) {
	list := []Todo{}
	DB.Find(&list)
	fmt.Printf("Users: %#v \n", list)
	data:= struct {
        Status string
        Data []Todo
        Vars map[string]string
    }{
    	Status: "success",
    	Data: list,
    	Vars: i.GetVars(r),
    }
	w.WriteHeader(http.StatusOK)
	i.Render(w, "index", data)
}
