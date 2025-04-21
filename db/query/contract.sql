-- name: CreateContract :one
INSERT INTO contracts (contract_type, price_cents)
VALUES ($1 , $2)
RETURNING *;

-- name: GetContract :one
SELECT * FROM contracts
WHERE id = $1
LIMIT 1;

-- name: ListContracts :many
SELECT * FROM contracts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteContract :exec
DELETE FROM contracts where id = $1;


