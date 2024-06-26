-- name: CheckUsernameAndEmail :one 
SELECT
  id,
  email <> '' AS email_present,
  username <> '' AS username_present
FROM
  users WHERE username = $1 or email = $2;

-- name: GetUserByName :one
SELECT * FROM users WHERE username = $1;

-- name: CreateProfile :one
INSERT INTO profiles (
  first_name,
  last_name,
  addresses_id,
  profile_image_url,
  phone_number,
  company_number,
  whatsapp_number,
  gender,
  all_languages_id,
  ref_no,
  cover_image_url
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: CreateUser :one
INSERT INTO users (
  email,
  username,
  hashed_password,
  status,
  roles_id,
  profiles_id,
  user_types_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;
