package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"openpowerlab/backend/internal/server"
)

const defaultPort = ":8080"

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := server.NewServer(defaultPort)
	if err := srv.Run(ctx); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}
