package handler

import (
	"net/http"
	"order_service/internal/model"
	"order_service/internal/service"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) AddItem(c *gin.Context) {
	var input model.CartItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect JSON format"})
		return
	}

	//todo: JWT
	input.UserID = 1

	err := h.service.AddItem(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item successfully added to cart"})
}

func (h *CartHandler) GetCart(c *gin.Context) {
	//temporary hardcode
	userID := int64(1)

	cart, err := h.service.GetCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve cart"})
		return
	}

	c.JSON(http.StatusOK, cart)
}
