package control

import "fmt"

func SwitchStuff() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("start of week")
	case "Friday":
		fmt.Println("weekend near")
	default:
		fmt.Println("normal day")
	}
}
