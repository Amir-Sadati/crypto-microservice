package asset

import (
	"context"

	assetpb "github.com/Amir-Sadati/crypto-microservice/crypto-proto/proto/asset"
)

type Service struct {
	assetClient assetpb.AssetServiceClient
}

func NewService(assetClient assetpb.AssetServiceClient) *Service {
	return &Service{
		assetClient: assetClient,
	}
}

func (s *Service) CreateAsset(c context.Context, request CreateAssetRequest) error {
	req := &assetpb.CreateAssetRequest{
		Symbol:   request.Symbol,
		Name:     request.Name,
		Type:     assetpb.AssetType(assetpb.AssetType_value[request.Type]),
		Decimals: request.Decimals,
		IsActive: request.IsActive,
	}
	_, err := s.assetClient.CreateAsset(c, req)
	return err
}
