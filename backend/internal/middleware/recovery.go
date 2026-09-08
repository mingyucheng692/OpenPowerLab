// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Recovery returns a middleware that recovers from panics, logs the error with
// a stack trace, and returns a 500 response.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.ErrorContext(r.Context(), "http panic recovered",
					"error", err,
					"path", r.URL.Path,
					"method", r.Method,
					"remote_addr", r.RemoteAddr,
					"stack", string(debug.Stack()),
				)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
