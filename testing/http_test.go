package mytests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http"))
	})
	
	handler.ServeHTTP(rr, req)
	
	if rr.Body.String() != "hello http" {
		t.Errorf("expected 'hello http', got %s", rr.Body.String())
	}
}
