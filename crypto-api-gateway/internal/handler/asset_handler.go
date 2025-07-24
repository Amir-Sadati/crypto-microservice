package handler

import (
	"net/http"
	"strings"

	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/response"
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/service/asset"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AssetHandler struct {
	assetService *asset.Service
	validator    *validator.Validate
}

func NewAssetHandler(assetService *asset.Service, validator *validator.Validate) *AssetHandler {
	return &AssetHandler{
		assetService: assetService,
		validator:    validator,
	}
}

// CreateAsset godoc
//
//	@Summary		Create a new asset
//	@Description	Creates an asset by receiving JSON data
//	@Tags			assets
//	@Accept			json
//	@Produce		json
//	@Param			request	body		asset.CreateAssetRequest	true	"body"
//	@Success		200		{object}	response.ApiResponseNoData
//	@Failure		400		{object}	response.ApiResponseNoData
//	@Failure		404		{object}	response.ApiResponseNoData
//	@Failure		500		{object}	response.ApiResponseNoData
//	@Router /api/v1/asset/create [post]
func (h *AssetHandler) CreateAsset(c *gin.Context) {
	var request asset.CreateAssetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.WriteFailNoData(c.Writer, http.StatusBadRequest, err.Error(), "")
		return
	}
	request.Type = strings.ToUpper(request.Type)

	err := h.assetService.CreateAsset(c.Request.Context(), request)
	if err != nil {
		response.WriteFailFromGrpc(c.Writer, err)
		return
	}

	response.WriteSuccessNoData(c.Writer, "Asset created successfully")
}
