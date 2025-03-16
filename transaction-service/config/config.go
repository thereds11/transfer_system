package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DBUrl       string
	KafkaBroker string
}

func LoadConfig() Config {
	_ = godotenv.Load()
	return Config{
		Port:        getEnv("PORT", "8081"),
		DBUrl:       getEnv("DB_URL", "postgres://transferuser:secret@localhost:5432/transferdb?sslmode=disable"),
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
