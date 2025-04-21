-- name: GetKyc :one
SELECT * FROM kyc_verification
WHERE user_id = $1 LIMIT 1;

-- name: ListKycs :many
SELECT * FROM kyc_verification ORDER BY user_id
LIMIT $1 OFFSET $2;

-- name: UpdateKycStatus :exec
UPDATE kyc_verification
SET kyc_status = $2 WHERE user_id = $1
RETURNING *;