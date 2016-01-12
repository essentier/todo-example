package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/essentier/todo-example/db"
	"gopkg.in/mgo.v2/bson"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB(r)
	var todos Todos
	db.C("todo").Find(bson.M{}).All(&todos)
	log.Printf("To dos: %#v", todos)

	uj, _ := json.Marshal(todos)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// Stub an user to be populated from the body
	log.Printf("CreateTodo called")
	u := Todo{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.Id = bson.NewObjectId()

	// Write the user to mongo
	db := db.GetDB(r)
	db.C("todo").Insert(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}
