-- name: CreateAccount :one
INSERT INTO authors (
  id,
  login,
  password
) VALUES (
  $1, $2, $3
) RETURNING *