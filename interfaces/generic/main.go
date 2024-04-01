package main

import (
	"context"
	"fmt"
)

// IEvent는 이벤트 인터페이스입니다.
type IEvent interface {
}

// IEventHandler는 제네릭으로 정의된 이벤트 핸들러 인터페이스입니다.
type IEventHandler[T IEvent, Tx any] func(ctx context.Context, tx Tx, event T) error

// 예제 이벤트 구조체
type ExampleEvent struct {
	Message string
}

// 예제 이벤트 핸들러
func exampleEventHandler(ctx context.Context, tx string, event ExampleEvent) error {
	fmt.Printf("Handling event: %v\n", event)
	fmt.Printf("Transaction details: %v\n", tx)
	// 여기에 실제 이벤트 처리 로직을 추가할 수 있습니다.
	return nil
}

func main() {
	// 예제 이벤트 생성
	event := ExampleEvent{
		Message: "Hello, World!",
	}

	// 예제 이벤트 핸들러 생성
	handler := IEventHandler[ExampleEvent, string](exampleEventHandler)

	// 예제 트랜잭션 정보
	transaction := "SampleTransaction123"

	// 이벤트 핸들러 호출
	err := handler(context.Background(), transaction, event)
	if err != nil {
		fmt.Printf("Error handling event: %v\n", err)
	}
}
