package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sqlx.DB
}

func Init(dbURL, migrationsDir string) *sqlx.DB {
	applyMigrations(dbURL, migrationsDir)

	conn, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Dtabase connection established")
	return conn
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
