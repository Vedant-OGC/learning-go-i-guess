package functions

import "fmt"

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func Closures() {
	add10 := makeAdder(10)
	fmt.Println(add10(5)) // 15
}
