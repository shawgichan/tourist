-- name: CreatePlace :one
INSERT INTO places (
    name,
    description,
    opening_hours,
    closing_hours,
    rating,
    ticket_category,
    ticket_price,
    location_id,
    place_type_id,
    created_at,
    updated_at,
    cover_image_url,
    profile_image_url,
    resturant_branch_id,
    preference_match
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15
) RETURNING id;
