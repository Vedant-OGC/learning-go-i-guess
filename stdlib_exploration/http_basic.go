package stdlibexploration

import (
	"encoding/json"
	"net/http"
)

func jsonHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "hello api"})
}

func HttpBasic() {
	http.HandleFunc("/api", jsonHandler)
}
