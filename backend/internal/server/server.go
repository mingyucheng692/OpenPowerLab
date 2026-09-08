// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

// Package server assembles routing and the middleware chain, and manages the
// HTTP server lifecycle including graceful shutdown.
package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"openpowerlab/backend/internal/handler"
	"openpowerlab/backend/internal/middleware"
)

// Server represents the HTTP server with graceful shutdown capability.
type Server struct {
	httpServer *http.Server
}

// NewServer initializes routes, middleware chain, and returns a Server instance.
func NewServer(addr string) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", handler.HealthHandler)
	mux.HandleFunc("GET /api/ping", handler.PingHandler)

	// Chain: Recovery -> LimitBody -> CORS -> Mux
	wrappedHandler := middleware.Recovery(middleware.LimitBody(middleware.CORS(mux)))

	return &Server{
		httpServer: &http.Server{
			Addr:              addr,
			Handler:           wrappedHandler,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}
}

// Run starts the HTTP server and blocks until the context is cancelled, then executes graceful shutdown.
func (s *Server) Run(ctx context.Context) error {
	serverErrCh := make(chan error, 1)
	go func() {
		slog.Info("backend server listening", "addr", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrCh <- err
		}
	}()

	select {
	case err := <-serverErrCh:
		return err
	case <-ctx.Done():
		slog.Info("shutting down server gracefully...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.httpServer.Shutdown(shutdownCtx)
	}
}
