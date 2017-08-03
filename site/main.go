package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/stephenafamo/what-to-do/models"
)

func main() {
	defer models.DB.Close()
	router := mux.NewRouter()
	http.Handle("/", routes(router))
	http.ListenAndServe(":80", nil)
}