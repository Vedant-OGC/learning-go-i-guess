package stdlibexploration

import (
	"fmt"
	"time"
)

func TimeStuff() {
	// parsing time
	layout := "2006-01-02"
	str := "2026-05-18"
	t, _ := time.Parse(layout, str)
	fmt.Println("parsed time:", t)
}
