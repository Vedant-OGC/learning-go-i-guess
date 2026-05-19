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
	str := `{"page": 2, "fruits": ["banana", "pear"]}`
	var r Response
	json.Unmarshal([]byte(str), &r)
	fmt.Println("page:", r.Page, "fruits:", r.Fruits)
}
