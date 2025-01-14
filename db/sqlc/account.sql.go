// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: account.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO accounts (
  owner, balance,currency
) VALUES (
  $1, $2 , $3
)
RETURNING id, owner, balance, currency, created_at
`

type CreateAuthorParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAuthor, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
