package routes

import (
	"net/http"
	"transaction-service/config"
	"transaction-service/internal/transaction"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/segmentio/kafka-go"
)

func SetupRoutes(cfg config.Config, db *pgxpool.Pool, producer *kafka.Writer) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/transactions", transaction.HandleCreate(db, producer)).Methods("POST")
	return r
}
