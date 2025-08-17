package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func New(dbURL, migrationsDir string) *DB {
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connection established")

	applyMigrations(dbURL, migrationsDir)

	return &DB{Conn: conn}
}

func (d *DB) Close() {
	if err := d.Conn.Close(); err != nil {
		log.Printf("Failed to close database: %v", err)
	} else {
		log.Println("Database connection closed")
	}
}

func applyMigrations(dbURL, migrationsDir string) {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsDir),
		dbURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
