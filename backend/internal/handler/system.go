// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

package handler

import "net/http"

// HealthResponse represents the payload returned by HealthHandler.
type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

// PingResponse represents the payload returned by PingHandler.
type PingResponse struct {
	Message string `json:"message"`
}

// HealthHandler handles GET /api/health requests.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, HealthResponse{
		Status:  "ok",
		Service: "openpowerlab-backend",
	})
}

// PingHandler handles GET /api/ping requests.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, PingResponse{
		Message: "pong",
	})
}
