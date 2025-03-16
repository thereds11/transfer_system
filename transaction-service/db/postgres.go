package db

import (
	"context"
	"log"
	"transaction-service/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitPostgres(cfg config.Config) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), cfg.DBUrl)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}
	return pool
}
