package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Subscriber 소켓 생성
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5555")
	subscriber.SetSubscribe("A 데몬 상태 질의")

	for {
		// 마스터 데몬으로부터 메시지 수신
		msg, _ := subscriber.Recv(0)
		fmt.Println("수신된 메시지: ", msg)
	}
}
