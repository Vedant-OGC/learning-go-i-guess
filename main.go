package main

import (
	"fmt"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/todo_cli"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/expense_tracker"
)

func main() {
	fmt.Println("launcher: todo and expense tracker integrated")
	var _ = todo_cli.TodoList{}
	var _ = expense_tracker.ExpenseBook{}
}
