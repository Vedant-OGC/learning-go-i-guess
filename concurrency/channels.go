package concurrency

import "fmt"

func Channels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // closing channel
	
	// range over channel terminates when closed
	for elem := range queue {
		fmt.Println("received:", elem)
	}
}
