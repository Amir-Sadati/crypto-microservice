package grpchandler

import (
	"context"

	"github.com/Amir-Sadati/crypto-microservice/crypto-order-service/internal/model"
	"github.com/Amir-Sadati/crypto-microservice/crypto-order-service/internal/service/order"
	orderpb "github.com/Amir-Sadati/crypto-microservice/crypto-proto/proto/order"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderGrpcHandler struct {
	orderService *order.Service
	orderpb.UnimplementedOrderServiceServer
}

func NewOrderGrpcHandler(orderService *order.Service) *OrderGrpcHandler {
	return &OrderGrpcHandler{
		orderService: orderService,
	}
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse user_id to UUID: %v", err)
	}
	orderType := model.OrderType(req.OrderType)
	if !orderType.IsValid() {
		return nil, status.Errorf(codes.InvalidArgument, "invalid order type: %s", req.OrderType)
	}
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse amount to decimal: %v", err)
	}
	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse price to decimal: %v", err)
	}

	createOrderRequest := order.CreateOrderRequest{
		UserID:    userID,
		AssetID:   req.AssetId,
		OrderType: orderType,
		Amount:    amount,
		Price:     price,
	}

	err = h.orderService.CreateOrder(ctx, createOrderRequest)
	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{}, nil
}
