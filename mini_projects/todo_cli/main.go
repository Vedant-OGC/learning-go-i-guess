package todo_cli

import (
	"fmt"
	"os"
)

func RunTodoCLI() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo <add|list>")
		return
	}
	
	action := os.Args[1]
	list, _ := Load("todo.json")
	
	switch action {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("please specify task text")
			return
		}
		list.Add(os.Args[2])
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
	}
}
