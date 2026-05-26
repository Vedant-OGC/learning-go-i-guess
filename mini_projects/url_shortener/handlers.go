package url_shortener

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Code string `json:"code"`
}

type Handlers struct {
	Shortener *URLShortener
}

func (h *Handlers) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	code := h.Shortener.Shorten(req.URL)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{Code: code})
}

func (h *Handlers) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" || code == "shorten" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	url, ok := h.Shortener.Resolve(code)
	if !ok {
		http.NotFound(w, r)
		return
	}
	
	http.Redirect(w, r, url, http.StatusFound)
}
