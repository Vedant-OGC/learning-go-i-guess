package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rect struct {
	W, H float64
}

func (r Rect) Area() float64 { return r.W * r.H }
func (r Rect) Perimeter() float64 { return 2 * (r.W + r.H) }

func printShapeDetails(s Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

func Interfaces() {
	c := Circle{Radius: 5}
	r := Rect{W: 3, H: 4}
	printShapeDetails(c)
	printShapeDetails(r)
}
