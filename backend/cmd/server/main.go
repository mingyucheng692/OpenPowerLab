package main

import (
	"log"

	"openpowerlab/backend/internal/server"
)

const defaultPort = ":8080"

func main() {
	srv := server.NewServer(defaultPort)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}
}
