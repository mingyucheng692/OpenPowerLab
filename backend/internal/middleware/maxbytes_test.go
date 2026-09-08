// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// stubReaderHandler reads the entire request body and reports the read error.
func stubReaderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := readAllBody(r)
		if err != nil {
			http.Error(w, "request body too large", http.StatusRequestEntityTooLarge)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func TestLimitBodyAllowsWithinLimit(t *testing.T) {
	body := bytes.Repeat([]byte("a"), 64<<10) // exactly 64KB
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/anything", bytes.NewReader(body))

	LimitBody(stubReaderHandler()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d (64KB must pass)", rec.Code, http.StatusOK)
	}
}

func TestLimitBodyRejectsOverLimit(t *testing.T) {
	body := bytes.Repeat([]byte("a"), (64<<10)+1) // one byte over
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/anything", bytes.NewReader(body))

	LimitBody(stubReaderHandler()).ServeHTTP(rec, req)

	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusRequestEntityTooLarge)
	}
	if !strings.Contains(rec.Body.String(), "too large") {
		t.Errorf("body = %q, want it to mention the size limit", rec.Body.String())
	}
}
