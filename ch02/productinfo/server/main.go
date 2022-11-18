package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/coldbeverage/gRPC/ch02/productinfo/server/ecommerce"
	"google.golang.org/grpc"
)

const (
	host = "127.0.0.1"
	port = "50051"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}