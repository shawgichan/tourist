-- name: CreateRestaurant :one
INSERT INTO restaurants (
  name) VALUES ( $1) RETURNING *;

-- name: GetRestaurant :one
SELECT * FROM restaurants WHERE id = $1 LIMIT 1;

-- name: ListRestaurants :many
SELECT * FROM restaurants LIMIT $1 OFFSET $2;

