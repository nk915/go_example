package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 2001, "The server port")
)

type server struct {
	pb.UnimplementedPingServer
}

func (s *server) SayHello(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v", in.GetBody())

	return &pb.PingResponse{Body: "Server to " + in.GetBody()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPingServer(s, &server{}) // client가 사용할 수 있도록 등록
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
