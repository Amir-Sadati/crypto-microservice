package order

import (
	"context"
	"log"
	"time"

	"github.com/Amir-Sadati/crypto-microservice/crypto-order-service/internal/model"
	"github.com/jmoiron/sqlx"
)

//TODO: add sentry and log

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateOrder(c context.Context, request CreateOrderRequest) error {

	order := &model.Order{
		UserID:    request.UserID,
		AssetID:   request.AssetID,
		Amount:    request.Amount,
		Price:     request.Price,
		Status:    model.StatusPending,
		Type:      request.OrderType,
		CreatedAt: time.Now().UTC(),
	}

	query := `
		INSERT INTO orders (
			asset_id, user_id, order_type, amount, price, status, created_at, updated_at
		) VALUES (
			:asset_id, :user_id, :order_type, :amount, :price, :status, :created_at, :updated_at
		) RETURNING id;
	`

	rows, err := s.db.NamedQuery(query, order)
	if err != nil {
		log.Printf("order insert failed: %v\n", err)
		return err
	}
	defer rows.Close()
	return nil
}
