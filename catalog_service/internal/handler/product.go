package handler

import (
	"catalog_service/internal/model"
	"catalog_service/internal/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// handler needs repo to ask for data from it
type ProductHandler struct {
	repo *repository.ProductRepository
}

// putting repo in handler
func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	//ask from repo the list of products
	products, err := h.repo.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error to get menu"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "item id to GET is incorrect"})
		return
	}

	product, err := h.repo.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Create(ctx *gin.Context) {

	var input model.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect JSON format"})
		return
	}

	createdProduct, err := h.repo.Create(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create new product"})
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)

}

func (h *ProductHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "item id to PATCH is incorrect"})
		return
	}

	var input model.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect JSON format"})
		return
	}

	err = h.repo.Update(ctx.Request.Context(), id, input)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "could not find product"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not UPDATE new product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product was successfully updated!"})

}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "item id to DELETE is incorrect"})
		return
	}

	err = h.repo.Delete(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "could not find product"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not DELETE product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product was successfully deleted"})

}
