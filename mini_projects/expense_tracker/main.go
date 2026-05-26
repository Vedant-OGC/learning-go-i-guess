package expense_tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func RunExpenseTracker() {
	if len(os.Args) < 2 {
		fmt.Println("usage: expense <summary|add|delete>")
		return
	}
	
	action := os.Args[1]
	
	var book ExpenseBook
	data, err := os.ReadFile("expenses.json")
	if err == nil {
		json.Unmarshal(data, &book)
	}
	
	switch action {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("usage: expense add <amount> <category_id> <note>")
			return
		}
		amount, _ := strconv.ParseFloat(os.Args[2], 64)
		catID, _ := strconv.Atoi(os.Args[3])
		note := os.Args[4]
		
		err := book.Add(amount, Category(catID), note)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		newData, _ := json.Marshal(book)
		os.WriteFile("expenses.json", newData, 0644)
		fmt.Println("added expense")
		
	case "summary":
		total := CalculateTotal(book.Expenses)
		fmt.Printf("Total expenses: $%.2f\n", total)
		for _, e := range book.Expenses {
			fmt.Printf("- [%s] $%.2f: %s\n", e.Category.String(), e.Amount, e.Note)
		}
		
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("usage: expense delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if book.Delete(id) {
			newData, _ := json.Marshal(book)
			os.WriteFile("expenses.json", newData, 0644)
			fmt.Println("deleted expense")
		} else {
			fmt.Println("expense not found")
		}
	}
}
