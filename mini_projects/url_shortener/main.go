package url_shortener

import (
	"fmt"
	"net/http"
)

func RunURLShortener() {
	shortener := NewShortener()
	handlers := &Handlers{Shortener: shortener}
	
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)
	
	fmt.Println("URL Shortener Server listening on :8081")
}
