package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/essentier/spickspan"
	"github.com/essentier/spickspan/config"
	"github.com/essentier/spickspan/model"
	"github.com/rs/cors"
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

func main() {
	n := negroni.Classic()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST"},
		AllowedHeaders:   []string{"Cache-Control", "Pragma", "Origin", "Authorization", "Content-Type", "Accept", "X-Requested-With"},
	})
	n.Use(c)

	provider, err := getServiceProvider()
	if err != nil {
		log.Fatalf("Could not resolve spickspan provider. The error is %v", err)
		return
	}

	mgoService, _ := provider.GetService("todo-rest-db")

	mgoUrl := mgoService.IP + ":" + strconv.Itoa(mgoService.Port)
	log.Printf("mgo url: %v", mgoUrl)
	n.Use(mongoMiddleware(mgoUrl, "tododb"))
	router := initRoutes()
	n.UseHandler(router)
	log.Printf("Listening on port 5000**bb")
	err = http.ListenAndServe(":5000", n)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Stopping...")
}
