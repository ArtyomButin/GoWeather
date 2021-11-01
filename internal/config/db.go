package config

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func Connection() {
	pool, err := pgxpool.Connect(context.Background(), "postgres://golang_user:golang_pass@go-db:5432/weather")
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer pool.Close()
}