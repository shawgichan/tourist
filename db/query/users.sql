-- name: CheckUsernameAndEmail :one 
SELECT
  id,
  email <> '' AS email_present,
  username <> '' AS username_present
FROM
  users WHERE username = $1 or email = $2;
