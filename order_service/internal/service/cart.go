package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"order_service/internal/model"
	"order_service/internal/pkg"
	"time"

	"go.uber.org/zap"
)

type CartRepository interface {
	AddItem(ctx context.Context, item model.CartItem) error
	GetByUserID(ctx context.Context, userID int64) ([]model.CartItem, error)
}

// dependency injection
type CartService struct {
	repo       CartRepository
	catalogURL string
	client     *http.Client
}

func NewCartService(repo CartRepository, catalogURL string) *CartService {
	return &CartService{
		repo:       repo,
		catalogURL: catalogURL,
		client:     &http.Client{Timeout: 5 * time.Second},
	}
}

func (s *CartService) AddItem(ctx context.Context, item model.CartItem) error {
	log := pkg.FromContext(ctx)

	if item.Quantity <= 0 {
		log.Warn("Attempted to add item with invalid quantity", zap.Int("quantity", item.Quantity))
		return errors.New("quantity must be greater than zero")
	}

	url := fmt.Sprintf("%s/api/v1/menu/%d", s.catalogURL, item.ProductID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Error("Failed to create request to catalog", zap.Error(err))
		return errors.New("internal error")
	}

	resp, err := s.client.Do(req)
	if err != nil {
		log.Error("Failed to reach catalog service", zap.Error(err), zap.String("url", url))
		return errors.New("catalog service is unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		log.Warn("Product not found in catalog", zap.Int("product_id", item.ProductID))
		return errors.New("product does not exist")
	}
	if resp.StatusCode != http.StatusOK {
		log.Error("Catalog service returned error", zap.Int("status", resp.StatusCode))
		return errors.New("failed to verify product")
	}

	log.Info("Product verified via catalog, adding to cart", zap.Int("product_id", item.ProductID))
	return s.repo.AddItem(ctx, item)
}

func (s *CartService) GetCart(ctx context.Context, userID int64) ([]model.CartItem, error) {
	return s.repo.GetByUserID(ctx, userID)
}
