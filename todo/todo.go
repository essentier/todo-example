package todo

// type Entity interface {
// 	SetId(id db.EntityId)
// 	GetId() db.EntityId
// }

type EntityId string

type Todo struct {
	Id        EntityId `json:"id" bson:"_id"`
	Name      string   `json:"name"`
	Completed bool     `json:"completed"`
}

// func (todo *Todo) SetId(id db.EntityId) {
// 	todo.Id = id
// }

// func (todo *Todo) GetId() db.EntityId {
// 	return todo.Id
// }

type Todos []Todo
