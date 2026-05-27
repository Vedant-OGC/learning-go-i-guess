package url_shortener

import (
	"fmt"
	"net/http"
)

func RunURLShortener() {
	shortener := NewShortener()
	handlers := &Handlers{Shortener: shortener}
	
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handlers.ShortenHandler)
	mux.HandleFunc("/", handlers.RedirectHandler)
	
	fmt.Println("URL Shortener Server listening on :8081 (Press Ctrl+C to stop in launcher)")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Printf("Server stopped: %v\n", err)
	}
}
