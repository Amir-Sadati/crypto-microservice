package grpchandler

import (
	"context"
	"errors"
	"strings"

	"github.com/Amir-Sadati/crypto-microservice/crypto-asset-service/internal/model"
	"github.com/Amir-Sadati/crypto-microservice/crypto-asset-service/internal/service/asset"
	assetpb "github.com/Amir-Sadati/crypto-microservice/crypto-proto/proto/asset"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AssetGrpcHandler struct {
	AssetService *asset.Service
	assetpb.UnimplementedAssetServiceServer
}

func NewAssetGrpcHandler(AssetService *asset.Service) *AssetGrpcHandler {
	return &AssetGrpcHandler{
		AssetService: AssetService,
	}
}

func (h *AssetGrpcHandler) CreateAsset(ctx context.Context, req *assetpb.CreateAssetRequest) (*assetpb.CreateAssetResponse, error) {
	AssetType := model.AssetType(strings.ToLower(req.Type.String()))
	if !AssetType.IsValid() {
		return nil, status.Errorf(codes.InvalidArgument, "invalid Asset type: %s", req.Type.String())
	}
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name cannot be empty")
	}
	if req.Symbol == "" {
		return nil, status.Errorf(codes.InvalidArgument, "symbol cannot be empty")
	}
	if req.Decimals < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "decimals cannot be negative")

	}
	createAssetRequest := asset.CreateAssetRequest{
		Symbol:   req.Symbol,
		Name:     req.Name,
		Type:     AssetType,
		Decimals: req.Decimals,
		IsActive: req.IsActive,
	}

	err := h.AssetService.CreateAsset(ctx, createAssetRequest)
	if errors.Is(err, asset.ErrAssetAlreadyExists) {
		return nil, status.Errorf(codes.AlreadyExists, "asset with symbol '%s' already exists", req.Symbol)
	}
	if err != nil {
		return nil, err
	}

	return &assetpb.CreateAssetResponse{}, nil
}
