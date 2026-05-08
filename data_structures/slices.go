package datastructures

import "fmt"

func Slices() {
	// slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:4] // s has elements 20, 30, 40
	fmt.Println("slice:", s)
	fmt.Println("len:", len(s), "cap:", cap(s))
}
