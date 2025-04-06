package kafka

import (
	"transaction-service/config"

	"github.com/segmentio/kafka-go"
)

type KafkaIO struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
}

func InitKafkaProducer(cfg config.Config) *KafkaIO {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.KafkaBroker},
		Topic:    "transaction.created",
		Balancer: &kafka.LeastBytes{},
	})

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   "transaction.fraud",
		GroupID: "fraud-consumer-1",
	})
	return &KafkaIO{
		Writer: writer,
		Reader: reader,
	}
}
