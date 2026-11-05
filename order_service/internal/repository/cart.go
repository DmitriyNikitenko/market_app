package repository

import (
	"context"
	"fmt"
	"order_service/internal/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddItem(ctx context.Context, item model.CartItem) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO cart_items (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, product_id) 
		DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity;
	`

	_, err := r.db.Exec(ctx, query, item.UserID, item.ProductID, item.Quantity)
	if err != nil {
		return fmt.Errorf("could not add item to cart: %w", err)
	}

	return nil
}

func (r *CartRepository) GetByUserID(ctx context.Context, userID int64) ([]model.CartItem, error) {
	query := `
		SELECT id, user_id, product_id, quantity 
		FROM cart_items 
		WHERE user_id = $1 
		ORDER BY id ASC;
	`

	//
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.CartItem
	for rows.Next() {
		var item model.CartItem
		if err := rows.Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
