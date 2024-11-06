package storer

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MySQLStorer struct {
	db *sqlx.DB
}

func NewMySQLStorer(db *sqlx.DB) *MySQLStorer {
	return &MySQLStorer{db: db}
}

func (ms *MySQLStorer) CreateProduct(ctx context.Context, p *Product) (*Product, error) {
	query := `INSERT INTO products(name, image, category, description, rating, num_reviews, price, count_in_stock) VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)`
	res, err := ms.db.NamedExecContext(ctx, query, p)
	if err != nil {
		return nil, fmt.Errorf("error inserting the product: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error inserting lastindexid : %v", err)
	}
	p.ID = id
	return p, nil
}

func (ms *MySQLStorer) GetProduct(ctx context.Context, id int64) (*Product, error) {
	var p Product
	query := `SELECT * FROM products WHERE id=?`
	err := ms.db.GetContext(ctx, &p, query, id)
	if err != nil {
		return nil, fmt.Errorf("error getting product: %v", err)
	}
	return &p, nil
}

func (ms *MySQLStorer) ListProducts(ctx context.Context) ([]*Product, error) {
	var products []*Product
	query := `SELECT * FROM products`
	err := ms.db.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, fmt.Errorf("error getting list of products: %v", err)
	}
	return products, nil
}

func (ms *MySQLStorer) UpdateProduct(ctx context.Context, p *Product) (*Product, error) {
	query := `UPDATE products SET name=:name, image=:image, category:category, description:description, rating:rating, num_reviews:num_reviews, price:price, count_in_stock:count_in_stock WHERE id=:id`
	_, err := ms.db.NamedExecContext(ctx, query, p)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %v", err)
	}
	return p, nil
}

func (ms *MySQLStorer) DeleteProduct(ctx context.Context, id int64) error {
	query := `DELETE FROM products WHERE id=?`
	_, err := ms.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf(" error deleting product: %v", err)
	}
	return nil
}
