package datastructures

import "fmt"

type Person struct {
	Name string
	Age  int
}

// receiver parameter (value receiver)
func (p Person) Greet() string {
	return "Hello " + p.Name
}

func Structs() {
	p := Person{Name: "Bob", Age: 20}
	fmt.Println(p.Greet())
}
