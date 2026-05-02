package basics

import "fmt"

const Pi = 3.14159

// iota auto increments!
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
)

func Constants() {
	fmt.Println(Sunday, Monday, Tuesday)
}
