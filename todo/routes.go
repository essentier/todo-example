package todo

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	router.Handle("/todos",
		negroni.New(
			negroni.Wrap(http.HandlerFunc(GetTodos)),
		)).Methods("GET")

	router.Handle("/todos",
		negroni.New(
			negroni.Wrap(http.HandlerFunc(CreateTodo)),
		)).Methods("POST")

}
