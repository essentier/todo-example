package todo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name"`
	Completed bool          `json:"completed"`
	Due       time.Time     `json:"due"`
}

type Todos []Todo
