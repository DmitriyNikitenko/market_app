package main

import (
	"catalog_service/internal/config"
	"catalog_service/internal/handler"
	"catalog_service/internal/repository"
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	//initializing config via NewConfigMust() method
	cfg := config.NewConfigMust()

	//creating context to listen OS signals
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Could not create connection pool with DataBase: %v", err)
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		log.Fatalf("DataBase ping is failed: %v", err)
	}
	log.Println("Connection to DataBase is SUCCESSFULL")

	//app layer assembling
	repo := repository.NewProductRepository(dbPool)
	productHandler := handler.NewProductHandler(repo)

	r := gin.Default()

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
		log.Printf("Server is running on port: %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server running error: %v", err)
		}
	}()

	//waiting for context decline
	<-ctx.Done()
	log.Println("Shutting down server gracefully...")

	//giving server some time(5sec) to finish current requests
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting successfully")

}
