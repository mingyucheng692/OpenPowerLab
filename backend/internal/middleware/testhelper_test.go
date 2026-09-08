// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

package middleware

import (
	"io"
	"net/http"
)

// readAllBody drains the request body and returns the read error, if any.
func readAllBody(r *http.Request) (int, error) {
	n, err := io.Copy(io.Discard, r.Body)
	return int(n), err
}
