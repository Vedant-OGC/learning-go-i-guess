package interfaces

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func Interfaces() {
	fmt.Println("interfaces intro")
}
