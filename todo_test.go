package main

import (
	"testing"

	"github.com/essentier/testutil"
)

func TestTodoRestAPI(t *testing.T) {
	t.Parallel()
	t.Logf("done")
	todoService := testutil.CreateRestService("todo-rest", t)
	defer todoService.Release() // This will also release the mongodb service.

	// hostUrl := "http://" + todoService.IP + ":" + strconv.Itoa(todoService.Port)
	// log.Printf("host url: %v", hostUrl)
	// serviceReady := spickspan.WaitService(todoService)
	// if serviceReady {
	// 	log.Printf("service is ready")
	// }

	// var createdTodo todo.Todo
	// newTodo := todo.Todo{Name: "todo1", Completed: false}
	// todoService.Resource("todos").Post(newTodo, &createdTodo)

	// var allTodos todo.Todos
	// todoService.Resource("todos").Get(&allTodos)
	// log.Printf("todos are: %#v", allTodos)
}

// func TestMongoDB(t *testing.T) {
// 	t.Parallel()
// 	mgoService, err := provider.GetService("mongodb")
// 	if err != nil {
// 		t.Errorf("Failed to get mongodb service. Error is: %#v. Error string is %v", err, err.Error())
// 	}

// 	url := mgoService.IP + ":" + strconv.Itoa(mgoService.Port)
// 	//url := "10.20.132.82:55351"
// 	log.Printf("mongodb url: %v", url)
// 	time.Sleep(100000 * time.Millisecond)
// 	log.Printf("after sleep")
// 	session, err := mgo.Dial(url)
// 	if err != nil {
// 		t.Errorf("Failed to connect to mongodb. Error is: %#v. Error string is %v", err, err.Error())
// 		return
// 	}

// 	reqSession := session.Clone()
// 	defer reqSession.Close()
// 	db := reqSession.DB("testdb")
// 	db.C("todo").Insert("hello")

// }

// var tries = 0

// func tryDialMock() bool {
// 	tries += 1
// 	if tries < 30 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

// func TestWait(t *testing.T) {
// 	waitService(spickspan.Service{})
// 	time.Sleep(5 * time.Second)
// }
