package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/todo_cli"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/expense_tracker"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/url_shortener"
)

func main() {
	fmt.Println("done learning go! welcome to the final interactive launcher.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n==================================")
		fmt.Println("    LEARNING GO I GUESS Launcher  ")
		fmt.Println("==================================")
		fmt.Println("1) Run TODO CLI Project")
		fmt.Println("2) Run Expense Tracker")
		fmt.Println("3) Run URL Shortener API Server")
		fmt.Println("q) Exit")
		fmt.Print("Choose an option: ")
		
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "q" {
			fmt.Println("bye!")
			break
		}
		
		switch input {
		case "1":
			fmt.Println("Launching TODO CLI...")
			todo_cli.RunTodoCLI()
		case "2":
			fmt.Println("Launching Expense Tracker...")
			expense_tracker.RunExpenseTracker()
		case "3":
			fmt.Println("Starting URL Shortener Server (Press Ctrl+C to stop)...")
			url_shortener.RunURLShortener()
		default:
			fmt.Println("invalid choice")
		}
	}
}
