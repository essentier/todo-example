package main

import (
	"github.com/essentier/todo-example/todo"
	"github.com/gorilla/mux"
)

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	todo.SetRoutes(router)
	return router
}
