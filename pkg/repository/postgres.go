package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Config struct {
	ConnStr string
}

func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), cfg.ConnStr)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
		return nil, err
	}

	return pool, nil
}
