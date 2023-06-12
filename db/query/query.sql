-- name: Users :one
INSERT INTO Users (
    login, 
    password
) VALUES (
    $1, $2
) RETURNING *;
