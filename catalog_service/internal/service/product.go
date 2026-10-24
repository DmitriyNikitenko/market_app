package service

import (
	"catalog_service/internal/model"
	"catalog_service/internal/pkg"
	"context"
	"errors"

	"go.uber.org/zap"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]model.Product, error)
	GetByID(ctx context.Context, id int) (model.Product, error)
	Create(ctx context.Context, product model.Product) (model.Product, error)
	Update(ctx context.Context, id int, product model.Product) error
	Delete(ctx context.Context, id int) error
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(ctx context.Context) ([]model.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProductService) GetByID(ctx context.Context, id int) (model.Product, error) {
	log := pkg.FromContext(ctx) //feat: added logger

	if id <= 0 {
		log.Warn("Attempted to GET product with invalid ID", zap.Float64("id", float64(id)))
		return model.Product{}, errors.New("invalid product id")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Create(ctx context.Context, p model.Product) (model.Product, error) {
	log := pkg.FromContext(ctx) //feat: added logger

	if p.Price <= 0 {
		log.Warn("Attempted to CREATE product with invalid price", zap.Float64("price", p.Price))
		return model.Product{}, errors.New("price must be greater than zero")
	}
	if p.Name == "" {
		log.Warn("Attempted to CREATE product with empty name")
		return model.Product{}, errors.New("product name cannot be empty")
	}
	return s.repo.Create(ctx, p)
}

func (s *ProductService) Update(ctx context.Context, id int, p model.Product) error {
	log := pkg.FromContext(ctx) //feat: added logger

	if id <= 0 {
		log.Warn("Attempted to UPDATE product with invalid id", zap.Float64("id", float64(p.ID)))
		return errors.New("invalid product id")
	}
	if p.Price <= 0 {
		log.Warn("Attempted to UPDATE product with invalid price", zap.Float64("price", p.Price))
		return errors.New("price must be greater than zero")
	}
	return s.repo.Update(ctx, id, p)
}

func (s *ProductService) Delete(ctx context.Context, id int) error {
	log := pkg.FromContext(ctx) //feat: added logger

	if id <= 0 {
		log.Warn("Attempted to DELETE product with invalid id", zap.Float64("id", float64(id)))
		return errors.New("invalid product id")
	}
	return s.repo.Delete(ctx, id)
}
