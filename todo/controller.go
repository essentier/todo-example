package todo

import (
	"encoding/json"
	"net/http"

	"github.com/essentier/nomockutil"
	"github.com/essentier/todo-example/db"
	"gopkg.in/mgo.v2/bson"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := getTodos(r)
	nomockutil.WriteObjectOrErr(w, todos, err)
}

func getTodos(r *http.Request) (Todos, error) {
	var todos Todos
	db, err := db.GetDB(r)
	if err != nil {
		return todos, err
	}

	err = db.C("todo").Find(bson.M{}).All(&todos)
	return todos, err
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := createTodo(r)
	nomockutil.WriteObjectOrErr(w, todo, err)
}

func createTodo(r *http.Request) (Todo, error) {
	todo := Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		return todo, err
	}

	db, err := db.GetDB(r)
	if err != nil {
		return todo, err
	}

	todo.Id = bson.NewObjectId()
	err = db.C("todo").Insert(todo)
	return todo, err
}
