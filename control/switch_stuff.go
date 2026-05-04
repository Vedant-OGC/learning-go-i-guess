package control

import "fmt"

func SwitchStuff() {
	// multiple cases
	switch num := 3; num {
	case 1, 3, 5:
		fmt.Println("odd")
	case 2, 4, 6:
		fmt.Println("even")
	}
}
