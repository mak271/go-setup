package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "my-topic",
		Balancer: &kafka.Hash{},
	})
	defer writer.Close()

	err := writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte("user-123"),
		Value: []byte("Hello, Kafka!"),
	})
	if err != nil {
		log.Fatal("Ошибка при отправке:", err)
	}
}
