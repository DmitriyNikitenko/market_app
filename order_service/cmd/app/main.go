package main

import (
	"context"
	"log"
	"order_service/internal/handler"
	"order_service/internal/pkg"
	"order_service/internal/repository"
	"order_service/internal/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {

	cfg := pkg.NewConfigMust()
	logger, err := pkg.NewLogger(cfg)
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}
	defer logger.Close()

	ctx := pkg.ToContext(context.Background(), logger)

	//подключение к бд заказов
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logger.Fatal("DATABASE_URL IS NOT SET")
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		logger.Fatal("Failed to connect to DB", zap.Error(err))
	}
	defer pool.Close()

	//внедрение dependency injection
	cartRepo := repository.NewCartRepository(pool)
	orderRepo := repository.NewOrderRepository(pool)

	//подключене к каталогу
	catalogURL := os.Getenv("CATALOG_SERVICE_URL")
	if catalogURL == "" {
		catalogURL = "http://catalog_service:8080"
	}

	//подключение к заказу
	orderURL := "" //заглушка

	cartService := service.NewCartService(cartRepo, catalogURL)
	cartHandler := handler.NewCartHandler(cartService)
	orderService := service.NewOrderService(orderRepo, orderURL)
	orderHandler := handler.NewOrderHandler(orderService)
	//настройка роутера Gin
	r := gin.Default()

	//группируем роуты
	api := r.Group("/api/v1")
	{
		api.POST("/cart", cartHandler.AddItem)
		api.GET("/cart", cartHandler.GetCart)
		api.POST("/checkout", orderHandler.Checkout)
	}

	//запускаем сервер
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Starting Order Service", zap.String("port", port))
	if err := r.Run(":" + port); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}

}
