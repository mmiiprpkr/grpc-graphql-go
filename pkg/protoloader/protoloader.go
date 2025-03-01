package protoloader

import (
	"grpcgqlgo/generated/proto/services"
	grpc "grpcgqlgo/pkg/GRPC"
	"log"
)

type Client struct {
	UserServiceClient    services.UserServiceClient
	ProductServiceClient services.ProductServiceClient
}

func CreateClient() Client {
	BackendGrpcServerServicesURL := "localhost:50051"

	conn, err := grpc.NewClient(BackendGrpcServerServicesURL)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	userServiceClient := services.NewUserServiceClient(conn)
	productServiceClient := services.NewProductServiceClient(conn)

	return Client{
		UserServiceClient:    userServiceClient,
		ProductServiceClient: productServiceClient,
	}
}
