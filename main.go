package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JulOuellet/sportlight/db"
)

func main() {
	const DB_URL = "postgres://postgres:password@localhost:5432/sportlight?sslmode=disable"

	database := db.New(DB_URL, "db/migrations")
	defer database.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
