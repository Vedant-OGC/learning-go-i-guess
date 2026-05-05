package functions

import "fmt"

func Funcs() {
	// using blank identifier to ignore values
	_, b := swap("hello", "world")
	fmt.Println("ignored first return:", b)
}

func swap(x, y string) (string, string) {
	return y, x
}
