package concurrency

import "fmt"

func Channels() {
	// deadlock demo
	// ch := make(chan int)
	// ch <- 1 // blocks forever because no receiver!
	
	fmt.Println("cannot write to unbuffered channel without active receiver")
}
