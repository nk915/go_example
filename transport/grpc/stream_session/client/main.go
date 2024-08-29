package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	pb "grpc/api"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:2001", "the address to connect to")
)

func main() {
	flag.Parse()

	// GRPC 서버에 연결
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// GRPC 클라이언트 생성
	client := pb.NewChatServiceClient(conn)

	// 클라이언 스트리밍을 위한 스트림 생성 (Server 커넥션)
	stream, err := client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	clientId := getClientId()
	done := make(chan bool)

	// 서버로부터 메시지를 수신하는 함수 호출
	go receiveMessages(stream, done)

	// 사용자 입력을 서버로 전송하는 함수 호출
	sendMessages(stream, clientId)

	<-done
	stream.CloseSend()

}

func getClientId() string {
	// 클라이언트 ID를 입력받음
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Client ID: ")
	clientID, _ := reader.ReadString('\n')
	clientID = strings.TrimSpace(clientID) // 개행 문자 및 공백 제거
	return clientID

}

func receiveMessages(stream pb.ChatService_ConnectClient, done chan bool) {
	defer close(done)

	// 서버로부터 메시지 스트림 응답 수신
	for {
		in, err := stream.Recv()
		// 서버가 스트림을 닫으면 종료
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Stream closed by server: %v", err)
			return
		}
		log.Printf("Server: %s\n", in.Message)

	}
}

func sendMessages(stream pb.ChatService_ConnectClient, clientId string) {
	scanner := bufio.NewScanner(os.Stdin)

	// 서버로 메시지 스트림 전송
	for scanner.Scan() {
		msg := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(msg, "done") {
			break
		}
		if err := stream.Send(&pb.ClientMessage{ClientId: clientId, Message: msg}); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
	}
}
