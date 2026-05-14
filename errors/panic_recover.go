package errorhandling

import "fmt"

func PanicRecover() {
	// defers execute in LIFO order
	defer fmt.Println("first defer (executes last)")
	defer fmt.Println("second defer (executes first)")
	
	fmt.Println("normal execution")
}
