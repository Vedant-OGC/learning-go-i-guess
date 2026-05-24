package url_shortener

import (
	"encoding/json"
	"net/http"
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
	var req ShortenRequest
	json.NewDecoder(r.Body).Decode(&req)
	
	code := h.Shortener.Shorten(req.URL)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{Code: code})
}
