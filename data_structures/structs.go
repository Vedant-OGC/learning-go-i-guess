package datastructures

import "fmt"

func Structs() {
	// anonymous struct
	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog.name, "is good boy?", dog.isGood)
}
