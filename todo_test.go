package main

import (
	"log"
	"testing"

	"github.com/essentier/spickspan"
	"github.com/essentier/testutil"
	"github.com/essentier/todo-example/todo"
)

// type failTestRestErrHanlder struct {
// 	t *testing.T
// }

// func (h *failTestRestErrHanlder) HandleError(err error, message string) {
// 	if err != nil {
// 		h.t.Fatalf(message+" Error is: %v.", err)
// 	}
// }

// func TestTodo2(t *testing.T) {
// 	errHandler := &failTestRestErrHanlder{t: t}
// 	service := model.Service{Protocol: "http", IP: "104.196.39.254", Port: 45802, Id: "test-todo-rest"}
// 	api := gopencils.Api(service.GetUrl())
// 	rw := &testutil.ResWrapper{Resource: api, ErrHandler: errHandler}
// 	todoService := &testutil.RestService{Service: service, Api: rw}

// 	var createdTodo todo.Todo
// 	//res := rw.NewChildResource("todos")
// 	newTodo := todo.Todo{Name: "todo5", Completed: false}
// 	todoService.Resource("todos").Post(newTodo, &createdTodo)
// 	//res.Post(newTodo, &createdTodo)
// 	// if err != nil {
// 	// 	t.Logf("post error is %v", err)
// 	// }
// 	//return &RestService{Service: service, api: rw}
// }

func TestTodoRestAPI(t *testing.T) {
	t.Parallel()
	t.Logf("done")
	todoService := testutil.CreateRestService("todo-rest", t)
	defer todoService.Release() // This will also release the mongodb service.

	log.Printf("host url: %v", todoService.Service.GetUrl())
	serviceReady := spickspan.ProbeService(todoService.Service, "/todos")
	if serviceReady {
		log.Printf("service is ready")
	}

	var createdTodo todo.Todo
	newTodo := todo.Todo{Name: "todo1", Completed: false}
	todoService.Resource("todos").Post(newTodo, &createdTodo)

	// var allTodos todo.Todos
	// todoService.Resource("todos").Get(&allTodos)
	// log.Printf("todos are: %#v", allTodos)
}

// func TestTodoRestAPI(t *testing.T) {
// 	t.Parallel()
// 	t.Logf("done")
// 	todoService := testutil.CreateRestService("todo-rest", t)
// 	defer todoService.Release() // This will also release the mongodb service.

// 	log.Printf("host url: %v", todoService.Service.GetUrl())
// 	serviceReady := spickspan.WaitService(todoService.Service)
// 	if serviceReady {
// 		log.Printf("service is ready")
// 	}

// 	service := model.Service{Protocol: "http", IP:"127.0.0.1", Port:5000, Id:"test-todo-"}
// 	api := gopencils.Api(service.GetUrl())

// 	var createdTodo todo.Todo
// 	newTodo := todo.Todo{Name: "todo1", Completed: false}
// 	todoService.Resource("todos").Post(newTodo, &createdTodo)

// 	var allTodos todo.Todos
// 	todoService.Resource("todos").Get(&allTodos)
// 	log.Printf("todos are: %#v", allTodos)
// }

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
