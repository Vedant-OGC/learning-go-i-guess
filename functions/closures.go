package functions

import "fmt"

func Closures() {
	// closure capture bug
	var funcs []func()
	for i := 0; i < 4; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i) // captured by reference!
		})
	}
	
	for _, f := range funcs {
		f()
	}
}
