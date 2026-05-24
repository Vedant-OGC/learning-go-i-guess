package url_shortener

import (
	"math/rand"
	"time"
)

type URLShortener struct {
	Urls map[string]string
}

func NewShortener() *URLShortener {
	return &URLShortener{Urls: make(map[string]string)}
}

func (s *URLShortener) Shorten(url string) string {
	code := generateCode()
	s.Urls[code] = url
	return code
}

func (s *URLShortener) Resolve(code string) (string, bool) {
	url, ok := s.Urls[code]
	return url, ok
}

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
