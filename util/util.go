package util

import (
	"net/http"

	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
)

type key int

const DB_KEY key = 0

func GetDB(r *http.Request) *mgo.Database {
	if rv := context.Get(r, DB_KEY); rv != nil {
		return rv.(*mgo.Database)
	}
	return nil
}
