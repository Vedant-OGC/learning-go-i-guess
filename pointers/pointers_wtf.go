package pointers

import "fmt"

func PointersWtf() {
	i := 42
	p := &i
	fmt.Println("address:", p)
	fmt.Println("value:", *p)
}
