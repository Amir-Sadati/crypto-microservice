package asset

import "github.com/Amir-Sadati/crypto-microservice/crypto-asset-service/internal/model"

type CreateAssetRequest struct {
	Symbol   string          `json:"symbol"`
	Name     string          `json:"name"`
	Type     model.AssetType `json:"type"`
	Decimals int32           `json:"decimals"`
	IsActive bool            `json:"is_active"`
}
