package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Checkout(ctx context.Context, userID uint64) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//начало транзакции
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx) //гарантия отката

	var orderID int

	err = tx.QueryRow(ctx, "INSERT INTO orders (user_id, status) VALUES ($1, 'created') RETURNING id", userID).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	//sql запрос для копирования товаров
	queryToCopyGoods := `
		INSERT INTO order_items (order_id, product_id, quantity) 
		SELECT $1, product_id, quantity
		FROM cart_items
		WHERE user_id = $2;
	`

	_, err = tx.Exec(ctx, queryToCopyGoods, orderID, userID)
	if err != nil {
		return 0, err
	}

	//sql запрос для очистки корзины
	queryToClear := `
		DELETE FROM cart_items
		WHERE user_id  = $1;
	`
	_, err = tx.Exec(ctx, queryToClear, userID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return int(orderID), nil
}
