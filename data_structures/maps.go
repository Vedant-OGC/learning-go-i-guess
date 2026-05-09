package datastructures

import "fmt"

func Maps() {
	m := map[string]int{"foo": 1, "bar": 2}
	delete(m, "foo")
	fmt.Println("after delete:", m)
}
