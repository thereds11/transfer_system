package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/segmentio/kafka-go"
)

func ListenForFraudEvents(reader *kafka.Reader, dbConn *pgxpool.Pool) {
	go func() {
		ctx := context.Background()
		for {
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Kafka read error: %v", err)
				continue
			}

			var tx map[string]interface{}
			if err := json.Unmarshal(msg.Value, &tx); err != nil {
				log.Printf("Failed to parse transaction: %v", err)
				continue
			}

			sourceID := tx["source_account_id"].(string)
			_, err = dbConn.Exec(context.Background(), `
				UPDATE transactions
				SET fraud_flag = true
				WHERE source_account_id = $1
			`, sourceID)
			if err != nil {
				log.Printf("DB update failed: %v", err)
				continue
			}
			log.Printf("🚨 Transaction flagged as fraudulent for account: %s", sourceID)
		}
	}()
}
