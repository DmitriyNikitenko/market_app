package handler

import (
	"net/http"
	"order_service/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	//получение контекста запроса
	ctx := c.Request.Context()
	userID := uint64(1)

	orderID, err := h.service.Checkout(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created", "order_id": orderID})

}
