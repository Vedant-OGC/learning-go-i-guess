package mytests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/Vedant-OGC/learning-go-i-guess/mini_projects/url_shortener"
)

func TestURLShortener(t *testing.T) {
	shortener := url_shortener.NewShortener()
	handlers := &url_shortener.Handlers{Shortener: shortener}
	
	reqBody, _ := json.Marshal(map[string]string{"url": "https://example.com"})
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(reqBody))
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ShortenHandler)
	handler.ServeHTTP(rr, req)
	
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}
}
