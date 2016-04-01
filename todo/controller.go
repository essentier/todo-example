package todo

import (
	"encoding/json"
	"net/http"

	"github.com/essentier/nomockutil"
	"github.com/essentier/todo-example/db"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	db, err := db.GetDB(r)
	if err != nil {
		nomockutil.WriteError(http.StatusInternalServerError, w, err)
		return
	}

	repo := getRepository(db)
	todos, err := repo.getAll()
	nomockutil.WriteObjectOrErr(w, todos, http.StatusInternalServerError, err)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	db, err := db.GetDB(r)
	if err != nil {
		nomockutil.WriteError(http.StatusInternalServerError, w, err)
		return
	}

	todo := Todo{}
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		nomockutil.WriteError(http.StatusBadRequest, w, err)
		return
	}

	repo := getRepository(db)
	savedTodo, err := repo.saveNew(todo)
	nomockutil.WriteObjectOrErr(w, savedTodo, http.StatusInternalServerError, err)
}

// func createTodo(r *http.Request) (Todo, error) {
// 	todo := Todo{}
// 	err := json.NewDecoder(r.Body).Decode(&todo)
// 	if err != nil {
// 		return todo, err
// 	}

// 	db, err := db.GetDB(r)
// 	if err != nil {
// 		return todo, err
// 	}

// 	todo.Id = bson.NewObjectId()
// 	err = db.C("todo").Insert(todo)
// 	return todo, err
// }
