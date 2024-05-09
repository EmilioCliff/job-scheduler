package main

import (
	"context"
	"log"

	"github.com/EmilioCliff/event-scheduler/api"
	db "github.com/EmilioCliff/event-scheduler/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	conn, err := pgxpool.New(context.Background(), "postgresql://root:secret@localhost:5432/job-scheduler?sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to db: %w", err)
	}

	store := db.New(conn)
	server := api.NewServer(store)
	log.Println("Starting server at port 8080")
	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("failed to start server: %w", err)
	}
}
