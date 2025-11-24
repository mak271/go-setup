package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
		GroupID: "my-groupID",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received: %s, %s\n", string(msg.Key), string(msg.Value))
	}

}
