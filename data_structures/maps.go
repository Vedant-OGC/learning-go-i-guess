package datastructures

import "fmt"

func Maps() {
	m := map[string]int{"foo": 1, "bar": 2}
	
	// check if key exists
	v, ok := m["baz"]
	fmt.Println("value:", v, "exists?:", ok)
}
