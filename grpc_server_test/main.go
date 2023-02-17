package main

import (
	"grpc_example/news"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Create a listener on TCP port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a grpc server instance
	grpcServer := grpc.NewServer()

	server := news.Server{}

	// Attach the chat service to the server
	news.Init()
	news.RegisterNewsServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
