package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderType string

const (
	BuyOrder  OrderType = "buy"
	SellOrder OrderType = "sell"
)

func (o OrderType) IsValid() bool {
	switch o {
	case BuyOrder, SellOrder:
		return true
	default:
		return false
	}
}

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusCompleted OrderStatus = "completed"
	StatusCancelled OrderStatus = "cancelled"
	StatusFailed    OrderStatus = "failed"
	StatusExpired   OrderStatus = "expired"
)

type Order struct {
	ID        int32           `db:"id"`
	UserID    uuid.UUID       `db:"user_id"`
	AssetID   int32           `db:"asset_id"`
	Type      OrderType       `db:"order_type"`
	Amount    decimal.Decimal `db:"amount"`
	Price     decimal.Decimal `db:"price"`
	Status    OrderStatus     `db:"status"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
}
