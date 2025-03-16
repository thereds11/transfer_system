package main

import (
	"log"
	"net/http"
	"transaction-service/config"
	"transaction-service/db"
	"transaction-service/kafka"
	"transaction-service/routes"
)

func main() {
	cfg := config.LoadConfig()
	dbConn := db.InitPostgres(cfg)
	producer := kafka.InitKafkaProducer(cfg)

	handler := routes.SetupRoutes(cfg, dbConn, producer)
	log.Printf("Starting server on port %s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, handler)
}
