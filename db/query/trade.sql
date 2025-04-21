-- name: CreateTrade :one
INSERT INTO trades (buy_order_id, sell_order_id, contract_id, price_cents, quantity)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetTrade :one
SELECT * FROM trades
WHERE id = $1
LIMIT 1;

-- name: ListTrades :many
SELECT * FROM trades
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteTrade :exec
DELETE FROM trades WHERE id = $1;