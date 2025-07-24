package order

type CreateOrderRequest struct {
	UserID    string `json:"userId"`
	AssetID   int32  `json:"assetId"`
	OrderType string `json:"orderType"`
	Amount    string `json:"amount"`
	Price     string `json:"price"`
}
