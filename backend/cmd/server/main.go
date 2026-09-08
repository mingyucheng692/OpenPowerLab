// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

// Package main is the entry point of the OpenPowerLab backend. It wires up
// logging, signal-based lifecycle management, and the HTTP server.
package main

import (
	"cmp"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"openpowerlab/backend/internal/server"
)

const defaultPort = "8080"

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	addr := ":" + cmp.Or(os.Getenv("PORT"), defaultPort)
	srv := server.NewServer(addr)
	if err := srv.Run(ctx); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}
