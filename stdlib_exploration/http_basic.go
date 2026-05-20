package stdlibexploration

import (
	"fmt"
	"net/http"
)

func HttpBasic() {
	// query params
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprintf(w, "User: %s", name)
	})
}
