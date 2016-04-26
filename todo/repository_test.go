package todo

import (
	"testing"

	"github.com/essentier/testutil"
	"github.com/essentier/todo-example/db"
)

func TestSaveNew(t *testing.T) {
	t.Parallel()
	mgoService := testutil.CreateMgoService("todo-db", t)
	defer mgoService.Release()
	dbSession, err := db.CreateDBSession(mgoService.GetUrl())
	handleFatalError(t, "Failed to create DB session.", err)
	defer dbSession.Close()

	repo := getRepository(dbSession.DB("tododb"))
	todo := Todo{Name: "todo1", Completed: false}
	savedTodo, err := repo.saveNew(todo)
	handleFatalError(t, "Failed to find todo by ID.", err)

	retrievedTodo, err := repo.findById(savedTodo.Id)
	handleFatalError(t, "Failed to find todo by ID.", err)
	if retrievedTodo.Name != "todo1" {
		t.Errorf("Expect todo's name to be todo1 but got %v", retrievedTodo.Name)
	}
}

func handleFatalError(t *testing.T, message string, err error) {
	if err != nil {
		t.Fatalf("%v Error is %v", message, err)
	}
}
