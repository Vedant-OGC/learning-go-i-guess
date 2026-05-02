package basics

import "fmt"

func Input() {
	var name string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}
