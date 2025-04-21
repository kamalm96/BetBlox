-- name: CreateOrder :one
INSERT INTO orders
(user_id, contract_id, order_type, order_style, price_cents, quantity) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1
LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET status = $1
WHERE id = $2;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;
