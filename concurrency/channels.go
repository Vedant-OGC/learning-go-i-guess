package concurrency

import "fmt"

func Channels() {
	// buffered channels allow sending values without receiver active
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
