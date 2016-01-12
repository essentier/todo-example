package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/codegangsta/negroni"
	"github.com/essentier/spickspan"
	"github.com/essentier/todo-example/db"
	"github.com/essentier/todo-example/todo"
)

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	todo.SetRoutes(router)
	return router
}

func main() {
	provider, err := spickspan.GetDefaultServiceProvider()
	if err != nil {
		log.Fatalf("Could not resolve spickspan provider. The error is %v", err)
		return
	}

	mgoService, err := spickspan.GetMongoDBService(provider, "todo-db")
	if err != nil {
		log.Fatalf("Could not get DB service. The error is %v", err)
		return
	}

	mgoUrl := mgoService.IP + ":" + strconv.Itoa(mgoService.Port)
	n := negroni.Classic()
	n.Use(db.MongoMiddleware(mgoUrl, "tododb"))
	router := initRoutes()
	n.UseHandler(router)
	log.Printf("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", n))
}
