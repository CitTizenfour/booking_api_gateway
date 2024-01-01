package services

import (
	"bk/bk_api_gateway/config"
	"bk/bk_api_gateway/genproto/order_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type ServiceManagerI interface {
	OrderService() order_service.OrderServiceClient
	TransportService() order_service.TransportServiceClient
	ReviewService() order_service.ReviewServiceClient
	ClienService() order_service.ClientServiceClient
}

type grpcClients struct {
	orderService     order_service.OrderServiceClient
	transportService order_service.TransportServiceClient
	reviewService    order_service.ReviewServiceClient
	clientService    order_service.ClientServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	resolver.SetDefaultScheme("dns")
	connOrderService, err := grpc.Dial(
		cfg.OrderServiceHost+":"+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		orderService:     order_service.NewOrderServiceClient(connOrderService),
		transportService: order_service.NewTransportServiceClient(connOrderService),
		reviewService:    order_service.NewReviewServiceClient(connOrderService),
		clientService:    order_service.NewClientServiceClient(connOrderService),
	}, nil
}

func (g *grpcClients) OrderService() order_service.OrderServiceClient {
	return g.orderService
}

func (g *grpcClients) TransportService() order_service.TransportServiceClient {
	return g.transportService
}

func (g *grpcClients) ReviewService() order_service.ReviewServiceClient {
	return g.reviewService
}

func (g *grpcClients) ClienService() order_service.ClientServiceClient {
	return g.clientService
}
