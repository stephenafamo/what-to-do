package main

import (
	"github.com/gorilla/mux"
)

func routes(router *mux.Router) *mux.Router {

	handler( router.Methods("GET").Subrouter(), "/", "IndexController", "Index")
	handler( router.Methods("POST").Subrouter(), "/todo", "TodoController", "Add")
	handler( router.Methods("PUT").Subrouter(), "/todo/{ID}", "TodoController", "Edit")
	handler( router.Methods("DELETE").Subrouter(), "/todo/{ID}", "TodoController", "Delete")
	handler( router.Methods("GET").Subrouter(), `/assets/{path:[a-zA-Z0-9=\-\/.]+}`, "AssetController", "Index")

	return router
}