package main

import (
	"context"
	"log/slog"
	"os"

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

	youtubeAPIKey := os.Getenv("YOUTUBE_API_KEY")
	if youtubeAPIKey != "" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		w := worker.New(database, youtubeAPIKey)
		w.Run(ctx)
		slog.Info("YouTube integration worker started")
	}

	e := web.RegisterRoutes(database)

	port := os.Getenv("PORT")
	if port == "" {
		slog.Error("PORT environment variable is not set")
		os.Exit(1)
	}

	if err := e.Start(":" + port); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
