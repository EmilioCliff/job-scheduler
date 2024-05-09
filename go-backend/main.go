package main

import (
	"context"
	"log"

	"github.com/EmilioCliff/event-scheduler/api"
	db "github.com/EmilioCliff/event-scheduler/db/sqlc"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	conn, err := pgxpool.New(context.Background(), "postgresql://root:secret@postgres1:5432/job-scheduler?sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to db: ", err)
	}

	m, err := migrate.New("file://db/migrations", "postgresql://root:secret@postgres1:5432/job-scheduler?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to load migration: ", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to run migrate up: ", err)
	}

	log.Println("Migration Successfull")

	store := db.New(conn)
	server := api.NewServer(store)
	log.Println("Starting server at port 8080")
	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("failed to start server: %w", err)
	}
}
