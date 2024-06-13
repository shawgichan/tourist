// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreatePlace(ctx context.Context, arg CreatePlaceParams) (int64, error)
	GetPlace(ctx context.Context, id int64) (Place, error)
	GetPlaces(ctx context.Context) ([]Place, error)
}

var _ Querier = (*Queries)(nil)
