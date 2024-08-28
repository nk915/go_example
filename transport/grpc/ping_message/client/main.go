package main

import (
	"context"
	"flag"
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

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPingClient(conn) // server의 method를 사용할 수 있도록 해줌

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.PingRequest{Body: "Client Body"})
	if err != nil {
		log.Fatalf("could not call: %v", err)
	}

	log.Printf("Call: %s", r.GetBody())

}
