package packages

import (
	"fmt"
	"os"
)

func FileIO() {
	err := os.WriteFile("test.txt", []byte("hello file"), 0644)
	if err != nil {
		panic(err)
	}
	
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	
	// cleanup
	os.Remove("test.txt")
}
