package basics

import (
	"bufio"
	"fmt"
	"os"
)

func Input() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your full name: ")
	if scanner.Scan() {
		name := scanner.Text()
		fmt.Println("Welcome, user:", name)
	}
}
