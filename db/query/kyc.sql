-- name: CreateKyc :one
INSERT INTO kyc_verification (user_id, ssn_last4, dob, address, kyc_status, submitted_at, verified_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetKyc :one
SELECT * FROM kyc_verification
WHERE user_id = $1 LIMIT 1;

-- name: UpdateKycStatus :exec
UPDATE kyc_verification
SET kyc_status = $2 WHERE user_id = $1
RETURNING *;