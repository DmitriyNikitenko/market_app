package handler

import (
	"catalog_service/internal/model"
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// handler needs repo to ask for data from it
type ProductService interface {
	GetAll(ctx context.Context) ([]model.Product, error)
	GetByID(ctx context.Context, id int) (model.Product, error)
	Create(ctx context.Context, product model.Product) (model.Product, error)
	Update(ctx context.Context, id int, product model.Product) error
	Delete(ctx context.Context, id int) error
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	//ask from service the list of products
	products, err := h.service.GetAll(ctx.Request.Context())
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

	product, err := h.service.GetByID(ctx.Request.Context(), id)
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

	createdProduct, err := h.service.Create(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	err = h.service.Update(ctx.Request.Context(), id, input)
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

	err = h.service.Delete(ctx.Request.Context(), id)
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

func (h *ProductHandler) UploadImage(ctx *gin.Context) {
	//извлекаем файл из http запроса по ключу "image"
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to get image"})
		return
	}

	//генерируем уникальное имя файла
	//filepath.Ext(file.Filename) достанет расширение(например .jpg или .png)
	extension := filepath.Ext(file.Filename)

	//создадим имя
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)

	//путь физического сохранения файла на диске сервера
	savePath := "./public/images/" + newFileName

	//сохранение файла на диск
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	//URL по которому будет доступна картинка
	imageURL := "/static/images/" + newFileName

	ctx.JSON(http.StatusOK, gin.H{"url": imageURL})
}
