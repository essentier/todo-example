package db

import (
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/go-errors/errors"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
)

type EntityId string

type key int

const (
	DB_KEY        key = 0
	dbDialTimeout     = 100 * time.Second
)

type mgoDBMiddleware struct {
	session  *mgo.Session
	database string
}

func (h *mgoDBMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	reqSession := h.session.Clone()
	defer reqSession.Close()
	db := reqSession.DB(h.database)
	context.Set(r, DB_KEY, db)
	next(rw, r)
}

func CreateDBMiddleware(url string, database string) (negroni.Handler, error) {
	session, err := CreateDBSession(url)
	if err != nil {
		return nil, err
	}
	return &mgoDBMiddleware{session: session, database: database}, nil
}

func CreateDBSession(url string) (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL(url)
	if err != nil {
		return nil, err
	}

	dialInfo.FailFast = false
	dialInfo.Timeout = dbDialTimeout
	return mgo.DialWithInfo(dialInfo)
}

func GetDB(r *http.Request) (*mgo.Database, error) {
	if rv := context.Get(r, DB_KEY); rv != nil {
		return rv.(*mgo.Database), nil
	} else {
		return nil, errors.Errorf("Database object for the request is not in the context .")
	}
}
