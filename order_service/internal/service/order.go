package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type OrderRepository interface {
	Checkout(ctx context.Context, userID uint64) (int, error)
}

// dependency injection
type OrderService struct {
	repo     OrderRepository
	orderURL string
	client   *http.Client
}

func NewOrderService(repo OrderRepository, orderURL string) *OrderService {
	return &OrderService{
		repo:     repo,
		orderURL: orderURL,
		client:   &http.Client{Timeout: 5 * time.Second},
	}
}

func (s *OrderService) Checkout(ctx context.Context, userID uint64) (int, error) {
	orderID, err := s.repo.Checkout(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("formatting error: %w", err)
	}
	return orderID, nil
}
