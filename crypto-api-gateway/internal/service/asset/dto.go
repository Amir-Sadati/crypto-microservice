package asset

type CreateAssetRequest struct {
	Symbol   string `json:"symbol" binding:"required"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Decimals int32  `json:"decimals"`
	IsActive bool   `json:"isActive"`
}
