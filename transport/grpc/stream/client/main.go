package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:2001", "the address to connect to")
)

func main() {
	flag.Parse()

	// GRPC 서버에 연결
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// GRPC 클라이언트 생성
	client := pb.NewStreamClient(conn)

	// 클라이언트 스트리밍을 위한 스트림 생성
	stream, err := client.StartStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// 서버로 메시지 스트림 전송
	go func() {
		messages := []string{"1", "2", "3", "4", "5"}
		for _, message := range messages {
			request := &pb.StreamMessage{Content: message}

			if err := stream.Send(request); err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
			time.Sleep(1 * time.Second)
		}

		// 스트림 종료
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error closing stream: %v", err)
		}
	}()

	// 서버로부터 메시지 스트림 응답 수신
	for {
		response, err := stream.Recv()
		// 서버가 스트림을 닫으면 종료
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}

		log.Printf("Received response from server: %s\n", response.GetContent())
	}

}
