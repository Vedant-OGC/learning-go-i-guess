package datastructures

import "fmt"

func Arrays() {
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	
	for i, v := range b {
		fmt.Printf("index %d = %d\n", i, v)
	}
}
