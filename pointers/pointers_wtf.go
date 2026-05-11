package pointers

import "fmt"

type Counter struct {
	val int
}

// value receiver doesn't mutate!
func (c Counter) AddValue() {
	c.val++
}

// pointer receiver mutates!
func (c *Counter) AddPointer() {
	c.val++
}

func PointersWtf() {
	c := Counter{val: 0}
	c.AddValue()
	fmt.Println("after value receiver:", c.val) // 0
	
	c.AddPointer()
	fmt.Println("after pointer receiver:", c.val) // 1
}
