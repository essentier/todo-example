package todo

import (
	"github.com/essentier/todo-example/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	getAll() (Todos, error)
	saveNew(todo Todo) (Todo, error)
	removeById(id db.EntityId) error
	findById(id db.EntityId) (Todo, error)
}

type repositoryImpl struct {
	c *mgo.Collection
}

func getRepository(db *mgo.Database) repository {
	return &repositoryImpl{c: db.C("todo")}
}

func (repo *repositoryImpl) getAll() (Todos, error) {
	var todos Todos
	err := repo.c.Find(bson.M{}).All(&todos)
	return todos, err
}

func (repo *repositoryImpl) removeById(id db.EntityId) error {
	return repo.c.RemoveId(id)
}

func (repo *repositoryImpl) saveNew(todo Todo) (Todo, error) {
	todo.Id = db.EntityId(bson.NewObjectId())
	err := repo.c.Insert(todo)
	return todo, err
}

func (repo *repositoryImpl) findById(id db.EntityId) (Todo, error) {
	var todo Todo
	err := repo.c.Find(bson.M{"_id": id}).One(&todo)
	return todo, err

}
