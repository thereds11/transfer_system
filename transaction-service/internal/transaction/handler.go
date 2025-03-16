package transaction

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"transaction-service/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/segmentio/kafka-go"
)

func HandleCreate(db *pgxpool.Pool, producer *kafka.Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tx Transaction
		cfg := config.LoadConfig()
		if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		_, err := db.Exec(context.Background(), `
			INSERT INTO transactions (source_account_id, target_account_id, amount)
			VALUES ($1, $2, $3)
		`, tx.SourceAccountID, tx.TargetAccountID, tx.Amount)
		if err != nil {
			log.Print(err)
			http.Error(w, "db insert error", http.StatusInternalServerError)
			return
		}
		data, _ := json.Marshal(tx)
		err = producer.WriteMessages(context.Background(), kafka.Message{Value: data})
		if err != nil {
			log.Print("kafka broker", cfg.KafkaBroker)
			log.Printf("❌ Failed to publish to Kafka: %v", err)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Transaction created"))
	}
}
