package stdlibexploration

import (
	"encoding/json"
	"fmt"
)

func JsonStuff() {
	// mapping to map[string]interface{}
	str := `{"name": "Jack", "metadata": {"active": true}}`
	var m map[string]interface{}
	json.Unmarshal([]byte(str), &m)
	fmt.Println(m["name"])
	
	meta := m["metadata"].(map[string]interface{})
	fmt.Println(meta["active"])
}
