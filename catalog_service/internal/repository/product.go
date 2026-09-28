package repository

import (
	"catalog_service/internal/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool //requests will be made from this pool
}

// accepts db connection and returns a ready to use repo
func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

// Method for sending SQL request
func (r *ProductRepository) GetAll(ctx context.Context) ([]model.Product, error) {

	query := `SELECT id, name, price, category, is_available FROM products ORDER BY id ASC`

	//Sending request to db
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close() //Closing the db answer in the end of function

	var products []model.Product //Massive with ready-products

	//Sorting out strigs that db gave after the function executes
	for rows.Next() {
		var p model.Product
		//Scan puts every item from db to Go structure
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category, &p.IsAvailable); err != nil {
			return nil, err
		}
		//Appending ready item to massive
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id int) (model.Product, error) {

	query := `
	SELECT id, name, price, category, is_available 
	FROM products
	WHERE id=$1;
	`

	row := r.db.QueryRow(ctx, query, id)

	var product model.Product

	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Category, &product.IsAvailable)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Product{}, fmt.Errorf(
				"product with id ='%d': %w",
				id,
				err,
			)
		}
		return model.Product{}, fmt.Errorf("scan error: %w", err)

	}

	return product, nil
}

func (r *ProductRepository) Create(ctx context.Context, product model.Product) (model.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO products (name, price, category, is_available) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`

	err := r.db.QueryRow(
		ctx,
		query,
		product.Name,
		product.Price,
		product.Category,
		product.IsAvailable,
	).Scan(&product.ID)

	if err != nil {
		return model.Product{}, fmt.Errorf("scan error: %w", err)
	}

	return product, err
}

func (r *ProductRepository) Update(ctx context.Context, id int, product model.Product) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		UPDATE products
		SET name=$1, price=$2, category=$3, is_available=$4
		WHERE id=$5;
	`

	cmdTag, err := r.db.Exec(ctx, query, product.Name, product.Price, product.Category, product.IsAvailable, id)

	if err != nil {
		return fmt.Errorf("Could not update product: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		DELETE FROM products
		WHERE id=$1;
	`

	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Could not delete product: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
