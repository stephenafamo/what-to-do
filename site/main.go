package main

import (
	"net/http"
	"github.com/stephenafamo/what-to-do/models"
)

func main() {
	defer models.DB.Close()
	router := customRouter()
	http.Handle("/", routes(router))
	http.ListenAndServe(":80", nil)
}