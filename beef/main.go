package main

import (
	"log"
	"net"

	beef "github.com/napakornsk/go-beef/proto"
	"google.golang.org/grpc"
)

type BeefServiceServer struct {
	beef.UnimplementedBeefServiceServer
}

func main() {
	server := grpc.NewServer()
	beef.RegisterBeefServiceServer(server, &BeefServiceServer{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Println("Starting gRPC server on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
