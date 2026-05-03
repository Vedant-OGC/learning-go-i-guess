package control

import "fmt"

func Conditionals() {
	x := 42
	if x > 10 {
		fmt.Println("greater than 10")
	} else {
		fmt.Println("smaller or equal")
	}
}
