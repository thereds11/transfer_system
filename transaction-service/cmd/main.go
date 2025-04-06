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
	kafkaIO := kafka.InitKafkaProducer(cfg)

	kafka.ListenForFraudEvents(kafkaIO.Reader, dbConn)

	handler := routes.SetupRoutes(cfg, dbConn, kafkaIO.Writer)
	log.Printf("Starting server on port %s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, handler)
}
