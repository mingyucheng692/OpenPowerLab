// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

package middleware

import "net/http"

// maxRequestBodySize caps the request body size to protect against memory DoS.
const maxRequestBodySize = 64 << 10 // 64KB

// LimitBody returns a middleware that wraps r.Body with http.MaxBytesReader.
// Exceeding the limit causes the request handlers to fail with a 413 response
// when they attempt to read the body.
func LimitBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)
		}
		next.ServeHTTP(w, r)
	})
}
