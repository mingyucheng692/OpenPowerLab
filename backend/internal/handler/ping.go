package handler

import (
	"encoding/json"
	"net/http"
)

// PingResponse represents the payload returned by PingHandler.
type PingResponse struct {
	Message string `json:"message"`
}

// PingHandler handles GET /api/ping requests.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	resp := PingResponse{
		Message: "pong",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
