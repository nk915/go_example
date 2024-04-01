package main

import (
	"fmt"
	"net"

	grpc "google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("failed to err: ", err)
	}

	grpcServer := grpc.NewServer()

	
}
