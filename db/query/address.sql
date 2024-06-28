-- name: CreateLocation :one
INSERT INTO locations (
    lat,
    lng,
    created_at,
    updated_at )
    VALUES ( $1, $2, $3, $4 ) RETURNING *;