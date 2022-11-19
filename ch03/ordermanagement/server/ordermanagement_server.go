package main

import (
	"context"

	pb "github.com/coldbeverage/gRPC/ch03/ordermanagement/server/ecommerce"
	wrappers "google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	orderMap map[string] *pb.Order
	pb.UnimplementedOrderManagementServer
}

func (s *server) GetOrder(ctx context.Context, orderId *wrappers.StringValue)(*pb.Order, error) {
	// Service Implementation
	order := orderId[orderId.Value]
	return &order, nil
}