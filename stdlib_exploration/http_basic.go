package stdlibexploration

import (
	"log"
	"net/http"
)

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request path:", r.URL.Path)
		next(w, r)
	}
}

func HttpBasic() {
	// middleware wrapping
	http.HandleFunc("/secure", loggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("secure area"))
	}))
}
