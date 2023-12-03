// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: products.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const CreateProduct = `-- name: CreateProduct :exec
INSERT INTO products (name, description, price, quantity)
VALUES ($1, $2, $3, $4)
`

type CreateProductParams struct {
	Name        string
	Description sql.NullString
	Price       sql.NullFloat64
	Quantity    int32
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.exec(ctx, q.createProductStmt, CreateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Quantity,
	)
	return err
}

const DeleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteProductStmt, DeleteProduct, id)
	return err
}

const GetProduct = `-- name: GetProduct :one
SELECT pr.id, pr.name, pr.description, pr.price, pr.quantity FROM products pr
WHERE pr.id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.queryRow(ctx, q.getProductStmt, GetProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Quantity,
	)
	return i, err
}

const GetProducts = `-- name: GetProducts :many
SELECT pr.id, pr.name, pr.description, pr.price, pr.quantity FROM products pr
ORDER BY pr.id
`

func (q *Queries) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.query(ctx, q.getProductsStmt, GetProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET name = $1, description = $2, price = $3, quantity = $4
WHERE id = $5
`

type UpdateProductParams struct {
	Name        string
	Description sql.NullString
	Price       sql.NullFloat64
	Quantity    int32
	ID          uuid.UUID
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.exec(ctx, q.updateProductStmt, UpdateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Quantity,
		arg.ID,
	)
	return err
}