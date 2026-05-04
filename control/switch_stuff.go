package control

import "fmt"

func SwitchStuff() {
	score := 85
	// switch without condition
	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 80:
		fmt.Println("Grade B") // python if-elif is longer than this!
	default:
		fmt.Println("Grade C")
	}
}
