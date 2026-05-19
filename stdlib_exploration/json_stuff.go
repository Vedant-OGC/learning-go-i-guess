package stdlibexploration

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func JsonStuff() {
	r := &Response{Page: 1, Fruits: []string{"apple", "peach"}}
	bolB, _ := json.Marshal(r)
	fmt.Println(string(bolB))
}
