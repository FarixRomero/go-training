// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Account, error)
}

var _ Querier = (*Queries)(nil)