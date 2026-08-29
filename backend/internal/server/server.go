package server

import (
	"fmt"
	"net/http"

	"openpowerlab/backend/internal/handler"
	"openpowerlab/backend/internal/middleware"
)

// Server represents the HTTP server.
type Server struct {
	addr    string
	handler http.Handler
}

// NewServer initializes routes, middleware, and returns a Server instance.
func NewServer(addr string) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", handler.HealthHandler)
	mux.HandleFunc("GET /api/ping", handler.PingHandler)

	wrappedHandler := middleware.CORS(mux)

	return &Server{
		addr:    addr,
		handler: wrappedHandler,
	}
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	fmt.Printf("Backend server listening on http://localhost%s\n", s.addr)
	return http.ListenAndServe(s.addr, s.handler)
}
