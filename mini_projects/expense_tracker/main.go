package expense_tracker

import (
	"fmt"
	"os"
)

func RunExpenseTracker() {
	if len(os.Args) < 2 {
		fmt.Println("usage: expense <summary|add>")
		return
	}
	fmt.Println("expense tracker started")
}
