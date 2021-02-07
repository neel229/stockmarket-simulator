// Code generated by sqlc. DO NOT EDIT.
// source: stock.sql

package db

import (
	"context"
	"database/sql"
)

const createStock = `-- name: CreateStock :one
INSERT INTO stock (
  ticker,
  creator_id,
  details
  ) VALUES (
  $1, $2, $3
) RETURNING id, ticker, creator_id, details
`

type CreateStockParams struct {
	Ticker    sql.NullString `json:"ticker"`
	CreatorID sql.NullInt64  `json:"creator_id"`
	Details   sql.NullString `json:"details"`
}

func (q *Queries) CreateStock(ctx context.Context, arg CreateStockParams) (Stock, error) {
	row := q.db.QueryRowContext(ctx, createStock, arg.Ticker, arg.CreatorID, arg.Details)
	var i Stock
	err := row.Scan(
		&i.ID,
		&i.Ticker,
		&i.CreatorID,
		&i.Details,
	)
	return i, err
}

const deleteStock = `-- name: DeleteStock :exec
DELETE FROM stock
WHERE id = $1
`

func (q *Queries) DeleteStock(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteStock, id)
	return err
}

const getStock = `-- name: GetStock :one
SELECT id, ticker, creator_id, details FROM stock
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetStock(ctx context.Context, id int64) (Stock, error) {
	row := q.db.QueryRowContext(ctx, getStock, id)
	var i Stock
	err := row.Scan(
		&i.ID,
		&i.Ticker,
		&i.CreatorID,
		&i.Details,
	)
	return i, err
}

const listStocks = `-- name: ListStocks :many
SELECT id, ticker, creator_id, details FROM stock
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListStocksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStocks(ctx context.Context, arg ListStocksParams) ([]Stock, error) {
	rows, err := q.db.QueryContext(ctx, listStocks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Stock
	for rows.Next() {
		var i Stock
		if err := rows.Scan(
			&i.ID,
			&i.Ticker,
			&i.CreatorID,
			&i.Details,
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

const updateStock = `-- name: UpdateStock :exec
UPDATE stock
SET details = $2 
WHERE id = $1
`

type UpdateStockParams struct {
	ID      int64          `json:"id"`
	Details sql.NullString `json:"details"`
}

func (q *Queries) UpdateStock(ctx context.Context, arg UpdateStockParams) error {
	_, err := q.db.ExecContext(ctx, updateStock, arg.ID, arg.Details)
	return err
}