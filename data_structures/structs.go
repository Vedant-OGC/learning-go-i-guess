package datastructures

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Structs() {
	p := Person{Name: "Bob", Age: 20}
	fmt.Println(p)
}
