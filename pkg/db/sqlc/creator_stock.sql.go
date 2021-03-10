// Code generated by sqlc. DO NOT EDIT.
// source: creator_stock.sql

package db

import (
	"context"
)

const createCreatorStock = `-- name: CreateCreatorStock :one
INSERT INTO creator_stock (creator_id, stock_id, mint_price, current_price)
VALUES ($1, $2, $3, $4)
RETURNING id, creator_id, stock_id, mint_price, current_price
`

type CreateCreatorStockParams struct {
	CreatorID    int64 `json:"creator_id"`
	StockID      int64 `json:"stock_id"`
	MintPrice    int32 `json:"mint_price"`
	CurrentPrice int32 `json:"current_price"`
}

func (q *Queries) CreateCreatorStock(ctx context.Context, arg CreateCreatorStockParams) (CreatorStock, error) {
	row := q.db.QueryRowContext(ctx, createCreatorStock,
		arg.CreatorID,
		arg.StockID,
		arg.MintPrice,
		arg.CurrentPrice,
	)
	var i CreatorStock
	err := row.Scan(
		&i.ID,
		&i.CreatorID,
		&i.StockID,
		&i.MintPrice,
		&i.CurrentPrice,
	)
	return i, err
}

const getCreatorStock = `-- name: GetCreatorStock :one
SELECT id, creator_id, stock_id, mint_price, current_price
FROM creator_stock
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCreatorStock(ctx context.Context, id int64) (CreatorStock, error) {
	row := q.db.QueryRowContext(ctx, getCreatorStock, id)
	var i CreatorStock
	err := row.Scan(
		&i.ID,
		&i.CreatorID,
		&i.StockID,
		&i.MintPrice,
		&i.CurrentPrice,
	)
	return i, err
}

const listCreatorStocks = `-- name: ListCreatorStocks :many
SELECT id, creator_id, stock_id, mint_price, current_price
FROM creator_stock
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListCreatorStocksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCreatorStocks(ctx context.Context, arg ListCreatorStocksParams) ([]CreatorStock, error) {
	rows, err := q.db.QueryContext(ctx, listCreatorStocks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreatorStock
	for rows.Next() {
		var i CreatorStock
		if err := rows.Scan(
			&i.ID,
			&i.CreatorID,
			&i.StockID,
			&i.MintPrice,
			&i.CurrentPrice,
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

const updateStockPrice = `-- name: UpdateStockPrice :exec
UPDATE creator_stock
SET current_price = $2
WHERE creator_id = $1
`

type UpdateStockPriceParams struct {
	CreatorID    int64 `json:"creator_id"`
	CurrentPrice int32 `json:"current_price"`
}

func (q *Queries) UpdateStockPrice(ctx context.Context, arg UpdateStockPriceParams) error {
	_, err := q.db.ExecContext(ctx, updateStockPrice, arg.CreatorID, arg.CurrentPrice)
	return err
}
