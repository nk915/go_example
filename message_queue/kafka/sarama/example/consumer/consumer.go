package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 카프카 브로커와 토픽, 그룹ID 설정
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"192.168.20.120:9092", "192.168.20.121:9092"},
		Topic:   "item_test",
		GroupID: "item_group",
	})

	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			fmt.Println("error while fetching message:", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		// 메시지를 성공적으로 수신했다고 컨슈머 그룹에 알림
		if err := r.CommitMessages(context.Background(), m); err != nil {
			fmt.Println("error while committing message:", err)
		}
	}

	if err := r.Close(); err != nil {
		fmt.Println("failed to close reader:", err)
	}
}
