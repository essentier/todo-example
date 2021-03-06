package main

import (
	"testing"

	"github.com/essentier/testutil"
	"github.com/essentier/todo-example/todo"
)

func TestTodoRestAPI(t *testing.T) {
	t.Parallel()
	todoService := testutil.CreateRestService("todo-example", "/todos", t)
	defer todoService.Release() // This will also release the mongodb service.

	var createdTodo todo.Todo
	newTodo := todo.Todo{Name: "todo1", Completed: false}
	todoService.Resource("todos").Post(newTodo, &createdTodo)

	var allTodos todo.Todos
	todoService.Resource("todos").Get(&allTodos)
	t.Logf("todos are: %#v", allTodos)
}
