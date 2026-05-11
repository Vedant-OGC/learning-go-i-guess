package pointers

import "fmt"

func zeroValue(x int) {
	x = 0
}

func zeroPointer(x *int) {
	*x = 0
}

func PointerFuncs() {
	i := 1
	zeroValue(i)
	fmt.Println("zeroValue:", i) // 1
	
	zeroPointer(&i)
	fmt.Println("zeroPointer:", i) // 0
}
