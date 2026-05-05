package functions

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Variadic() {
	fmt.Println(sum(1, 2, 3, 4))
}
