package main

import (
	"fmt"
	"github.com/Vedant-OGC/learning-go-i-guess/packages"
)

func main() {
	fmt.Println("testing my packages!")
	fmt.Println("5 + 3 =", packages.Add(5, 3))
	fmt.Println("reverse 'hello':", packages.Reverse("hello"))
}
