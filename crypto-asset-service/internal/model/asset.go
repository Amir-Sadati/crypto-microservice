package model

import "time"

type AssetType string

const (
	Coin  AssetType = "coin"
	Token AssetType = "token"
	Fiat  AssetType = "fiat"
)

func (o AssetType) IsValid() bool {
	switch o {
	case Coin, Token, Fiat:
		return true
	default:
		return false
	}
}

type Asset struct {
	ID        int32     `db:"id"`
	Symbol    string    `db:"symbol"`
	Name      string    `db:"name"`
	Type      AssetType `db:"type"`
	Decimals  int32     `db:"decimals"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
