package basics

import "fmt"

func Types() {
	var a int = 42
	var b float64 = 3.14
	var c string = "go is cool"
	var d bool = true
	
	fmt.Println(a, b, c, d)
	
	var e int = int(b)
	fmt.Println("float to int:", e)
	
	// string to int conversion?
	// val := int("123") // doesn't work like Python!
	// string(65) gives "A", which is weird
	fmt.Println("string conversion:", string(rune(65)))
}
