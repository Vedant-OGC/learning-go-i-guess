package interfaces

import "fmt"

func TypeAssertion() {
	var i interface{} = "hello"
	s, ok := i.(string)
	fmt.Println(s, ok)
	
	f, ok := i.(float64)
	fmt.Println(f, ok) // safe, doesn't panic
}
