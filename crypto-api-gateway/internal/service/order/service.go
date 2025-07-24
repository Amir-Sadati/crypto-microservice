package order

import (
	"context"

	orderpb "github.com/Amir-Sadati/crypto-microservice/crypto-proto/proto/order"
)

type Service struct {
	orderClient orderpb.OrderServiceClient
}

func NewService(orderClient orderpb.OrderServiceClient) *Service {
	return &Service{
		orderClient: orderClient,
	}
}

func (s *Service) CreateOrder(c context.Context, request CreateOrderRequest) error {
	req := &orderpb.CreateOrderRequest{
		UserId:    request.UserID,
		AssetId:   request.AssetID,
		OrderType: request.OrderType,
		Amount:    request.Amount,
		Price:     request.Price,
	}
	_, err := s.orderClient.CreateOrder(c, req)
	return err
}
