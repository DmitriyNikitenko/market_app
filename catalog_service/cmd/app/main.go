package main

import (
	"catalog_service/internal/handler"
	"catalog_service/internal/repository"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5433/menu_db?sslmode=disable"

	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer dbPool.Close()

	repo := repository.NewProductRepository(dbPool)

	productHandler := handler.NewProductHandler(repo)

	r := gin.Default()

	r.GET("/api/v1/menu", productHandler.GetMenu) //testing output of all products

	r.GET("/api/v1/menu/:id", productHandler.GetByID) //testing output of product by id

	r.POST("/api/v1/menu", productHandler.Create) //testing POST function for creating products

	r.PUT("/api/v1/menu/:id", productHandler.Update) //testing PUT function for updating products

	r.DELETE("/api/v1/menu/:id", productHandler.Delete) //testing DELETE function for deleting products

	r.Run(":8080")
}
