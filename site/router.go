package main

import (
	"github.com/gorilla/mux"
	"github.com/stephenafamo/what-to-do/controllers"
	"net/http"
	"reflect"
)

type stephenRouter interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
}

func handler(c stephenRouter, path string, controllerName string, method string) {
	theController := reflect.New(controller.Get(controllerName))
	theMethod := theController.MethodByName(method)
	c.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		theMethod.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r), theController.Elem().FieldByName("Params")})
	})
}
