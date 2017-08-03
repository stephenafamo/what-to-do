package main

import (
	"github.com/gorilla/mux"
)

func routes(router *mux.Router) *mux.Router {

	handler( router, "/", "IndexController", "Index")
	handler( router, "/add-todo", "TodoController", "Add")
	handler( router, `/assets/{path:[a-zA-Z0-9=\-\/.]+}`, "AssetController", "Index")

	return router
}