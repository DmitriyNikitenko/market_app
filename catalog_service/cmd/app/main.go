package main

import (
	"catalog_service/internal/config"
	"catalog_service/internal/handler"
	"catalog_service/internal/pkg"
	"catalog_service/internal/repository"
	"catalog_service/internal/service"
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {

	logCfg := pkg.NewConfigMust()
	logger, err := pkg.NewLogger(logCfg)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	//initializing config via NewConfigMust() method
	cfg := config.NewConfigMust()

	//creating context to listen OS signals
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Could not create connection pool with DataBase", zap.Error(err))
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		logger.Fatal("DataBase ping is failed", zap.Error(err))
	}
	logger.Info("Connection to DataBase is SUCCESSFULL")

	//app layer assembling
	repo := repository.NewProductRepository(dbPool)
	productService := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(productService)

	r := gin.Default()

	r.Use(handler.LoggerMiddleware(logger))

	api := r.Group("/api/v1")
	{
		api.POST("/menu", productHandler.Create)
		api.GET("/menu", productHandler.GetAll)
		api.GET("/menu/:id", productHandler.GetByID)
		api.PUT("/menu/:id", productHandler.Update)
		api.DELETE("/menu/:id", productHandler.Delete)
	}

	//packing Gin into HTTP-Server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	//starting server
	go func() {
		logger.Info("Server is running on port", zap.String("port", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("Server running error", zap.Error(err))
		}
	}()

	//waiting for context decline
	<-ctx.Done()
	logger.Info("Shutting down server gracefully...")

	//giving server some time(5sec) to finish current requests
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}
	logger.Info("Server exiting successfully")

}
