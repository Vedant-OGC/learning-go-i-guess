package pointers

import "fmt"

// lightbulb moment: pointers prevent copying big structs!
type BigStruct struct {
	data [1000]int
}

func updateStruct(b *BigStruct) {
	b.data[0] = 99
}

func PointerFuncs() {
	var b BigStruct
	updateStruct(&b)
	fmt.Println("updated struct without copying entire array:", b.data[0])
}
