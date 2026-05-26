package todo_cli

import (
	"fmt"
	"os"
	"strconv"
)

func RunTodoCLI() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo <add|list|complete>")
		return
	}
	
	action := os.Args[1]
	list, err := Load("todo.json")
	if err != nil {
		fmt.Println("Error loading:", err)
		return
	}
	
	switch action {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("please specify task text")
			return
		}
		err := list.Add(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		Save("todo.json", list)
		fmt.Println("added task")
	case "list":
		for _, t := range list.Todos {
			status := " "
			if t.Completed {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Text)
		}
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("please specify task ID")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if list.Complete(id) {
			Save("todo.json", list)
			fmt.Println("marked as complete")
		} else {
			fmt.Println("task not found")
		}
	}
}
