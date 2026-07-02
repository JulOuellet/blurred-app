package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JulOuellet/blurred-app/internal/db"
	"github.com/JulOuellet/blurred-app/internal/web"
	"github.com/JulOuellet/blurred-app/internal/worker"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		slog.Error("DATABASE_URL environment variable is not set")
		os.Exit(1)
	}

	migrationPath := "internal/db/migrations"

	database := db.Init(dbUrl, migrationPath)
	defer database.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	youtubeAPIKey := os.Getenv("YOUTUBE_API_KEY")
	if youtubeAPIKey != "" {
		w := worker.New(database, youtubeAPIKey)
		w.Run(ctx)
		slog.Info("YouTube integration worker started")
	}

	e := web.RegisterRoutes(database)
	e.Server.ReadHeaderTimeout = 5 * time.Second
	e.Server.ReadTimeout = 10 * time.Second
	e.Server.WriteTimeout = 30 * time.Second
	e.Server.IdleTimeout = 90 * time.Second

	port := os.Getenv("PORT")
	if port == "" {
		slog.Error("PORT environment variable is not set")
		os.Exit(1)
	}

	go func() {
		if err := e.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}
}
