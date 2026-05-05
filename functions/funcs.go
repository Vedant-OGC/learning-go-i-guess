package functions

import "fmt"

// multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return, dangerous but cool
}

func Funcs() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
