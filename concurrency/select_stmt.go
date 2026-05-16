package concurrency

import (
	"fmt"
)

func SelectStmt() {
	messages := make(chan string)
	
	// non-blocking read using default
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
}
