package url_shortener

import (
	"crypto/rand"
	"math/big"
)

type URLShortener struct {
	Urls map[string]string
}

func NewShortener() *URLShortener {
	return &URLShortener{Urls: make(map[string]string)}
}

func (s *URLShortener) Shorten(url string) string {
	code := generateSecureCode()
	s.Urls[code] = url
	return code
}

func (s *URLShortener) Resolve(code string) (string, bool) {
	url, ok := s.Urls[code]
	return url, ok
}

func generateSecureCode() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[n.Int64()]
	}
	return string(b)
}
