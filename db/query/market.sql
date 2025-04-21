-- name: CreateMarket :one
INSERT INTO markets
(title, description, category, status, closes_at, resolves_at)
VALUES ( $1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetMarket :one
SELECT * FROM markets
WHERE id = $1
LIMIT 1;

-- name: ListAllMarkets :many
SELECT * FROM markets
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListOpenMarkets :many
SELECT * FROM markets
WHERE status = 'open'
LIMIT $1
OFFSET $2;

-- name: ResolveMarket :exec
UPDATE markets
set status = $1
WHERE id = $2;


-- name: DeleteMarket :exec
DELETE FROM markets WHERE id = $1;
