package order

import (
	"github.com/Amir-Sadati/crypto-microservice/crypto-order-service/internal/model"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateOrderRequest struct {
	UserID    uuid.UUID       `json:"user_id"`
	AssetID   int32           `json:"asset_id"`
	OrderType model.OrderType `json:"order_type"`
	Amount    decimal.Decimal `json:"amount"`
	Price     decimal.Decimal `json:"price"`
}
