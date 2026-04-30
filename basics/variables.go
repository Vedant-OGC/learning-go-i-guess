package basics

import "fmt"

func Variables() {
	x := 10
	var y string = "hello"
	var z float64 = 20.5
	fmt.Println(x, y, z)
	
	// multiple declarations
	a, b := 1, 2
	fmt.Println(a, b)
}
