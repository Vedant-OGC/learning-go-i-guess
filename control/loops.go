package control

import "fmt"

func Loops() {
	// go only has for!
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}
}
