package datastructures

import "fmt"

type Address struct {
	City, State string
}

type Person struct {
	Name    string
	Age     int
	Address Address // nested
}

func Structs() {
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:  "Seattle",
			State: "WA",
		},
	}
	fmt.Println(p.Address.City)
}
