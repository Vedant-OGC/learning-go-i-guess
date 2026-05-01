package basics

import "fmt"

func Variables() {
	x := 10
	var y string = "hello"
	var z float64 = 20.5
	fmt.Println(x, y, z)
	
	a, b := 1, 2
	fmt.Println(a, b)
	
	// fixed the unused variable error!
	var used int = 5
	fmt.Println("used variable:", used)
}
