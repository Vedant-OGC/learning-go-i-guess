package control

import "fmt"

func Loops() {
	// for as while
	count := 0
	for count < 3 {
		fmt.Println("while equivalent count:", count)
		count++
	}
}
