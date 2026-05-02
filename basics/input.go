package basics

import (
	"bufio"
	"fmt"
	"os"
)

func Input() {
	// bufio scanner allows spaces in inputs
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter full name: ")
	if scanner.Scan() {
		name := scanner.Text()
		fmt.Println("Hello,", name)
	}
}
