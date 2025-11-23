package main

import (
	"log"
	"os"

	"github.com/JulOuellet/blurred-app/internal/db"
	"github.com/JulOuellet/blurred-app/internal/web"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	migrationPath := "internal/db/migrations"

	database := db.Init(dbUrl, migrationPath)
	defer database.Close()

	e := web.RegisterRoutes(database)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	log.Fatal(e.Start(":" + port))
}
