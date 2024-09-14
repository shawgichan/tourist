-- name: CreateRestaurantBranch :one
INSERT INTO restaurant_branches (
  name, resturant_id) VALUES ( $1, $2 ) RETURNING *;

-- name: GetRestaurantBranch :one
SELECT * FROM restaurant_branches WHERE id = $1 LIMIT 1;

-- name: ListRestaurantBranches :many
SELECT * FROM restaurant_branches WHERE resturant_id = $1 LIMIT $2 OFFSET $3;