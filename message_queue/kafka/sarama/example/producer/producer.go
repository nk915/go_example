package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 카프카 브로커 주소와 토픽 설정
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"192.168.20.120:9092", "192.168.20.121:9092"},
		Topic:   "item_test",
	})

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: "peter-test02",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
	)

	if err != nil {
		fmt.Println("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		fmt.Println("failed to close writer:", err)
	}
}
