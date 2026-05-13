package interfaces

import "fmt"

// empty interface{} is like dynamic typing in Python (or 'any' in new Go)
func TypeAssertion() {
	var x interface{} = 42
	fmt.Println("value of any type:", x)
}
