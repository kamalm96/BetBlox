-- name: CreateWallet :one
INSERT INTO wallets (user_id, balance_cents, locked_cents, updated_at)
VALUES ($1,$2,$3,$4) RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets
WHERE user_id = $1 LIMIT 1;

-- name: UpdateWallet :one
UPDATE wallets
SET balance_cents = $2
WHERE user_id = $1
RETURNING *;

-- name: UpdateLocked :one
UPDATE wallets
SET locked_cents = $2
WHERE user_id = $1
RETURNING *;

-- name: LogAudit :exec
INSERT INTO audit_logs (user_id, action, metadata, ip_address)
VALUES ($1, $2, $3, $4)
RETURNING *;



