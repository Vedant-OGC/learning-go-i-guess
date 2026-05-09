package datastructures

import "fmt"

func Slices() {
	original := []int{100, 200, 300}
	
	// use copy to prevent mutating original
	copied := make([]int, len(original))
	copy(copied, original)
	copied[0] = 999
	fmt.Println("original (safe):", original[0]) // still 100
}
