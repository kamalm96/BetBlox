-- name: CreateTransaction :one
INSERT INTO transactions (id, user_id, type, amount_cents, balance_after)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetTransactionById :one
SELECT * FROM transactions
WHERE id = $1
LIMIT 1;

-- name: GetAllTransactions :many
SELECT * FROM transactions
order by id
LIMIT $1
OFFSET $2;