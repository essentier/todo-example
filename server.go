package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-errors/errors"
	"github.com/gorilla/mux"

	"github.com/codegangsta/negroni"
	"github.com/essentier/spickspan"
	"github.com/essentier/spickspan/config"
	"github.com/essentier/spickspan/model"
	"github.com/essentier/spickspan/probe"
	"github.com/essentier/todo-example/db"
	"github.com/essentier/todo-example/todo"
)

func getServiceProvider() (model.Provider, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	registry, err := spickspan.GetDefaultKubeRegistry(config)
	if err != nil {
		return nil, err
	}

	return registry.ResolveProvider()
}

func getDBService(provider model.Provider) (model.Service, error) {
	mgoService, err := provider.GetService("todo-rest-db")
	if err != nil {
		return mgoService, err
	}

	serviceReady := probe.ProbeMgoService(mgoService)
	if serviceReady {
		return mgoService, nil
	} else {
		return mgoService, errors.Errorf("Service is not ready yet. The service is %v", mgoService)
	}
}

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	todo.SetRoutes(router)
	return router
}

func main() {
	provider, err := getServiceProvider()
	if err != nil {
		log.Fatalf("Could not resolve spickspan provider. The error is %v", err)
		return
	}

	mgoService, err := getDBService(provider)
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
