package main

import (
	"fmt"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/todo_cli"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/expense_tracker"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/url_shortener"
)

func main() {
	fmt.Println("--- launcher menu help ---")
	fmt.Println("1) Todo CLI")
	fmt.Println("2) Expense Tracker")
	fmt.Println("3) URL Shortener")
	var _ = todo_cli.TodoList{}
	var _ = expense_tracker.ExpenseBook{}
	var _ = url_shortener.URLShortener{}
}
