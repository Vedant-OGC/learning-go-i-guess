package datastructures

import "fmt"

func Slices() {
	s := []int{1, 2, 3, 4, 5}
	
	// delete element at index 2 (value 3)
	s = append(s[:2], s[3:]...)
	fmt.Println("after delete:", s) // [1 2 4 5]
	
	// slices are references!
	original := []int{100, 200, 300}
	copied := original
	copied[0] = 999
	fmt.Println("original:", original[0]) // mutated! cursed
}
