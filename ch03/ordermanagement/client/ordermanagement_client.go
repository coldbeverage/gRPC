package main

import (
	"context"
	"log"
	"time"

	pb "github.com/coldbeverage/gRPC/ch03/ordermanagement/client/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/wrappers"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	order, err := c.GetOrder(ctx, &wrappers.StringValue{Value: "106"})
	if err != nil {
		log.Fatalf("could not get order: %v", err)
	}
	log.Printf("Order: ", order.String())
}