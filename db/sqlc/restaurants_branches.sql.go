// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: restaurants_branches.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRestaurantBranch = `-- name: CreateRestaurantBranch :one
INSERT INTO restaurant_branches (
  name, resturant_id) VALUES ( $1, $2) RETURNING id, name, resturant_id
`

type CreateRestaurantBranchParams struct {
	Name        string      `json:"name"`
	ResturantID pgtype.Int8 `json:"resturant_id"`
}

func (q *Queries) CreateRestaurantBranch(ctx context.Context, arg CreateRestaurantBranchParams) (RestaurantBranch, error) {
	row := q.db.QueryRow(ctx, createRestaurantBranch, arg.Name, arg.ResturantID)
	var i RestaurantBranch
	err := row.Scan(&i.ID, &i.Name, &i.ResturantID)
	return i, err
}

const getRestaurantBranch = `-- name: GetRestaurantBranch :one
SELECT id, name, resturant_id FROM restaurant_branches WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRestaurantBranch(ctx context.Context, id int64) (RestaurantBranch, error) {
	row := q.db.QueryRow(ctx, getRestaurantBranch, id)
	var i RestaurantBranch
	err := row.Scan(&i.ID, &i.Name, &i.ResturantID)
	return i, err
}

const listRestaurantBranches = `-- name: ListRestaurantBranches :many
SELECT id, name, resturant_id FROM restaurant_branches WHERE resturant_id = $1 LIMIT $2 OFFSET $3
`

type ListRestaurantBranchesParams struct {
	ResturantID pgtype.Int8 `json:"resturant_id"`
	Limit       int32       `json:"limit"`
	Offset      int32       `json:"offset"`
}

func (q *Queries) ListRestaurantBranches(ctx context.Context, arg ListRestaurantBranchesParams) ([]RestaurantBranch, error) {
	rows, err := q.db.Query(ctx, listRestaurantBranches, arg.ResturantID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RestaurantBranch
	for rows.Next() {
		var i RestaurantBranch
		if err := rows.Scan(&i.ID, &i.Name, &i.ResturantID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
