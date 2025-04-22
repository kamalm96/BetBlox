-- name: GetResolution :one
SELECT * FROM market_resolution
WHERE market_id = $1
LIMIT 1;

-- name: ListResolutions :many
SELECT * FROM market_resolution
ORDER BY market_id
LIMIT $1
OFFSET $2;

-- name: MarkAsResolved :one
INSERT INTO market_resolution (market_id, outcome, resolved_by)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateResolution :one
UPDATE market_resolution
SET outcome = $1,
    resolved_at = $2,
    resolved_by = NOW()
WHERE market_id = $3
RETURNING *;

-- name: IsMarketResolved :one
SELECT EXISTS (
    SELECT 1 FROM market_resolution WHERE market_id = $1
) AS is_resolved;
