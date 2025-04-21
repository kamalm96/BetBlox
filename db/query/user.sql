-- name: CreateAccount :one
INSERT INTO users (email, username, password_hash)
VALUES ($1, $2, $3)
RETURNING id, email, username, created_at;

-- name: GetUser :one
SELECT email, username, created_at, is_verified FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT email, username, created_at, is_verified FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $1
WHERE email = $2;



