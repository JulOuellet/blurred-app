package main

import (
	"context"
	"log"
	"os"

	"github.com/JulOuellet/blurred-app/internal/db"
	"github.com/JulOuellet/blurred-app/internal/web"
	"github.com/JulOuellet/blurred-app/internal/worker"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
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
		log.Println("YouTube integration worker started")
	}

	e := web.RegisterRoutes(database)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	log.Fatal(e.Start(":" + port))
}
