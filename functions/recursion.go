package functions

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func Recursion() {
	fmt.Println(fib(10))
}
