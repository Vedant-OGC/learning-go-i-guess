package functions

import "fmt"

func Closures() {
	var funcs []func()
	for i := 0; i < 4; i++ {
		iCopy := i
		funcs = append(funcs, func() {
			fmt.Println(iCopy)
		})
	}
	
	for _, f := range funcs {
		f()
	}
}
