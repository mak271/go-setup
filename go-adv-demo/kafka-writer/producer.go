package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "my-topic",
		RequiredAcks: kafka.RequireAll,
		//Async:        true,
		Balancer: &kafka.Hash{},
	}
	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("Ошибка при отправке:", err)
	}
}
