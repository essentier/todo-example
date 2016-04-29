package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/codegangsta/negroni"
	"github.com/essentier/negronimgo"
	"github.com/essentier/spickspan"
	"github.com/essentier/todo-example/todo"
)

const (
	dbName    = "tododb"
	dbService = "todo-db"
)

func main() {
	provider, err := spickspan.GetDefaultServiceProvider()
	handleError("Failed to get spickspan provider.", err)

	mgoService, err := spickspan.GetMongoDBService(provider, dbService)
	handleError("Failed to get MongoDB service.", err)

	mgoUrl := mgoService.IP + ":" + strconv.Itoa(mgoService.Port)
	dbMiddleware, err := negronimgo.CreateDBMiddleware(mgoUrl, dbName)
	handleError("Failed to create DB middleware.", err)

	n := negroni.Classic()
	n.Use(dbMiddleware)
	router := initRoutes()
	n.UseHandler(router)
	log.Printf("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", n))
}

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	todo.SetRoutes(router)
	return router
}

func handleError(msg string, err error) {
	if err != nil {
		log.Fatalf("%v The error is %v", msg, err)
	}
}
