package main

import (
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/essentier/todo-example/util"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
)

//an example url is 127.0.0.1:27017
func mongoMiddleware(url string, database string) negroni.HandlerFunc {
	dialInfo, err := mgo.ParseURL(url)
	if err != nil {
		panic(err)
	}

	dialInfo.FailFast = false
	dialInfo.Timeout = 100 * time.Second

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	}

	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		reqSession := session.Clone()
		defer reqSession.Close()
		db := reqSession.DB(database)
		setDb(r, db)
		next(rw, r)
	})
}

func setDb(r *http.Request, val *mgo.Database) {
	context.Set(r, util.DB_KEY, val)
}
