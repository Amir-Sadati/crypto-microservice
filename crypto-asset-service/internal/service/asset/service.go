package asset

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Amir-Sadati/crypto-microservice/crypto-asset-service/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

// TODO: add sentry and log
var (
	ErrAssetAlreadyExists = errors.New("asset already exists")
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateAsset(c context.Context, request CreateAssetRequest) error {

	asset := &model.Asset{
		Symbol:    request.Symbol,
		Name:      request.Name,
		Type:      request.Type,
		Decimals:  request.Decimals,
		IsActive:  request.IsActive,
		CreatedAt: time.Now().UTC(),
	}

	query := `
		INSERT INTO assets (
			symbol, name, type, decimals, is_active, created_at, updated_at
		) VALUES (
			:symbol, :name, :type, :decimals, :is_active, :created_at, :updated_at
		) RETURNING id;
	`

	rows, err := s.db.NamedQuery(query, asset)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				// 23505 = unique_violation
				return ErrAssetAlreadyExists
			}
		}

		log.Printf("asset insert failed: %v\n", err)
		return err
	}

	defer rows.Close()
	return nil
}
