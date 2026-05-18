package stdlibexploration

import (
	"fmt"
	"time"
)

func TimeStuff() {
	t := time.Now()
	// why on earth is the formatting template 2006-01-02 15:04:05???
	// apparently it corresponds to 1 2 3 4 5 6 MST
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
