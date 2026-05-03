package control

import "fmt"

func Loops() {
	count := 0
	for count < 3 {
		fmt.Println("while equivalent count:", count)
		count++
	}
	
	// infinite loop with break
	i := 0
	for {
		if i == 5 {
			break
		}
		i++
	}
	// careful lol
}
