package functions

import "fmt"

func add(x int, y int) int {
	return x + y
}

func Funcs() {
	fmt.Println("add:", add(5, 7))
}
