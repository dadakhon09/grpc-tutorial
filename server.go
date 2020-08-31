package main

import (
	"fmt"
	"google.golang.org/grpc"
	"github.com/dadakhon09/grpc-tutorial/chat"
	"log"
	"net"
)


func main() {
	fmt.Println("Working")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer,  &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000: %v", err)
	}

}
