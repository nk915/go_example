package main

import (
	"flag"
	"fmt"
	pb "grpc/api"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 2001, "The server port")
)

type streamServer struct {
	pb.UnimplementedStreamServer
}

func (s *streamServer) StartStreaming(stream pb.Stream_StartStreamingServer) error {
	for {
		// 클라이언트로부터 메시지 스트림을 받음
		message, err := stream.Recv()
		if err == io.EOF {
			// 클라이언트가 스트림을 닫으면 종료
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received message from client: %s\n", message.GetContent())

		// 서버에서 응답을 생성하여 클라이언트에게 보냄
		response := &pb.StreamMessage{
			Content: "Server received message: " + message.GetContent(),
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}

}

func main() {
	flag.Parse()
	// GRPC 서버 리스너 생성
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// GRPC 서버 생성
	grpcServer := grpc.NewServer()

	// Stream 서비스 등록
	pb.RegisterStreamServer(grpcServer, &streamServer{})

	// 서버 시작
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
