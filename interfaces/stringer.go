package interfaces

import "fmt"

type Person struct {
	Name string
	Age  int
}

// implementing fmt.Stringer
func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func Stringer() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p) // calls String() implicitly
}
