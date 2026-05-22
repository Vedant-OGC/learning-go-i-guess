package mytests

import (
	"testing"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/todo_cli"
)

func TestTodoAdd(t *testing.T) {
	list := &todo_cli.TodoList{}
	list.Add("test task")
	if len(list.Todos) != 1 {
		t.Errorf("expected 1 task, got %d", len(list.Todos))
	}
	if list.Todos[0].Text != "test task" {
		t.Errorf("expected task text 'test task', got %s", list.Todos[0].Text)
	}
}
