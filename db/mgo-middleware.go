package db

import (
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
)

type key int

const DB_KEY key = 0

//an example url is 127.0.0.1:27017
func MongoMiddleware(url string, database string) negroni.HandlerFunc {
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
	context.Set(r, DB_KEY, val)
}

func GetDB(r *http.Request) *mgo.Database {
	if rv := context.Get(r, DB_KEY); rv != nil {
		return rv.(*mgo.Database)
	}
	return nil
}
