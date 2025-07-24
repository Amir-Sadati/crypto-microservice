package handler

import (
	"net/http"

	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/response"
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/service/order"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderHandler struct {
	orderService *order.Service
	validator    *validator.Validate
}

func NewOrderHandler(orderService *order.Service, validator *validator.Validate) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		validator:    validator,
	}
}

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Creates an order by receiving JSON data
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			request	body		order.CreateOrderRequest	true	"body"
//	@Success		200		{object}	response.ApiResponseNoData
//	@Failure		400		{object}	response.ApiResponseNoData
//	@Failure		404		{object}	response.ApiResponseNoData
//	@Failure		500		{object}	response.ApiResponseNoData
//	@Router /api/v1/order/create [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var request order.CreateOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.WriteFailNoData(c.Writer, http.StatusBadRequest, "Invalid request data", "")
		return
	}

	err := h.orderService.CreateOrder(c.Request.Context(), request)
	if err != nil {
		response.WriteFailFromGrpc(c.Writer, err)
		return
	}

	response.WriteSuccessNoData(c.Writer, "Order created successfully")
}
