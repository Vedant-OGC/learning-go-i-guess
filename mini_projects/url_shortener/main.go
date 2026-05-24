package url_shortener

import (
	"fmt"
)

func RunURLShortener() {
	shortener := NewShortener()
	handlers := &Handlers{Shortener: shortener}
	
	fmt.Println("URL Shortener handlers initialized", handlers)
}
