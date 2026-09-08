// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

// Package handler defines the JSON API endpoints of the OpenPowerLab backend.
package handler

import (
	"encoding/json"
	"net/http"
)

// writeJSON writes a JSON response with the specified status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
