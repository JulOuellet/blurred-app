package main

import (
	"log"

	"github.com/JulOuellet/sportlight/internal/db"
	"github.com/JulOuellet/sportlight/internal/web"
)

func main() {
	const DB_URL = "postgres://postgres:password@localhost:5432/sportlight?sslmode=disable"

	database := db.New(DB_URL, "internal/db/migrations")
	defer database.Close()

	e := web.RegisterRoutes(database)

	log.Println("Server running on localhost:8080")
	log.Fatal(e.Start(":8080"))
}
