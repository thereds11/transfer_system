package kafka

import (
	"transaction-service/config"

	"github.com/segmentio/kafka-go"
)

func InitKafkaProducer(cfg config.Config) *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.KafkaBroker},
		Topic:    "transaction.created",
		Balancer: &kafka.LeastBytes{},
	})
	return writer
}
