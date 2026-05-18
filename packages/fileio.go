package packages

import (
	"bufio"
	"os"
)

func FileIO() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	w := bufio.NewWriter(f)
	w.WriteString("buffered write\n")
	w.Flush()
	
	os.Remove("test.txt")
}
