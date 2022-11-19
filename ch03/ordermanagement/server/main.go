package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/coldbeverage/gRPC/ch03/ordermanagement/server/ecommerce"
	"google.golang.org/grpc"
)

const (
	host = "127.0.0.1"
	port = "50052"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})

	log.Printf("Starting gRPC listener on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}