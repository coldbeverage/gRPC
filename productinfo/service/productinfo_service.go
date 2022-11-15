package service

import (
	"context"
	pb "productinfo/service/ecommerce"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string] *pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product)(*pb.ProductID, error){
	out, err := uuid.New()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while generating ProductID", err)
	}
	return nil, nil
}