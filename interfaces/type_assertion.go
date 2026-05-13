package interfaces

import "fmt"

func TypeAssertion() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
}
