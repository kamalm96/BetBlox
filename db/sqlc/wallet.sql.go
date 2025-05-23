// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: wallet.sql

package db

import (
	"context"
	"database/sql"

	"github.com/sqlc-dev/pqtype"
)

const createWallet = `-- name: CreateWallet :one
INSERT INTO wallets (user_id, balance_cents, locked_cents, updated_at)
VALUES ($1,$2,$3,$4) RETURNING user_id, balance_cents, locked_cents, updated_at
`

type CreateWalletParams struct {
	UserID       int64        `json:"user_id"`
	BalanceCents int64        `json:"balance_cents"`
	LockedCents  int64        `json:"locked_cents"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, createWallet,
		arg.UserID,
		arg.BalanceCents,
		arg.LockedCents,
		arg.UpdatedAt,
	)
	var i Wallet
	err := row.Scan(
		&i.UserID,
		&i.BalanceCents,
		&i.LockedCents,
		&i.UpdatedAt,
	)
	return i, err
}

const getWallet = `-- name: GetWallet :one
SELECT user_id, balance_cents, locked_cents, updated_at FROM wallets
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetWallet(ctx context.Context, userID int64) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, getWallet, userID)
	var i Wallet
	err := row.Scan(
		&i.UserID,
		&i.BalanceCents,
		&i.LockedCents,
		&i.UpdatedAt,
	)
	return i, err
}

const logAudit = `-- name: LogAudit :exec
INSERT INTO audit_logs (user_id, action, metadata, ip_address)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, action, metadata, ip_address, created_at
`

type LogAuditParams struct {
	UserID    sql.NullInt64         `json:"user_id"`
	Action    string                `json:"action"`
	Metadata  pqtype.NullRawMessage `json:"metadata"`
	IpAddress pqtype.Inet           `json:"ip_address"`
}

func (q *Queries) LogAudit(ctx context.Context, arg LogAuditParams) error {
	_, err := q.db.ExecContext(ctx, logAudit,
		arg.UserID,
		arg.Action,
		arg.Metadata,
		arg.IpAddress,
	)
	return err
}

const updateLocked = `-- name: UpdateLocked :one
UPDATE wallets
SET locked_cents = $2
WHERE user_id = $1
RETURNING user_id, balance_cents, locked_cents, updated_at
`

type UpdateLockedParams struct {
	UserID      int64 `json:"user_id"`
	LockedCents int64 `json:"locked_cents"`
}

func (q *Queries) UpdateLocked(ctx context.Context, arg UpdateLockedParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, updateLocked, arg.UserID, arg.LockedCents)
	var i Wallet
	err := row.Scan(
		&i.UserID,
		&i.BalanceCents,
		&i.LockedCents,
		&i.UpdatedAt,
	)
	return i, err
}

const updateWallet = `-- name: UpdateWallet :one
UPDATE wallets
SET balance_cents = $2
WHERE user_id = $1
RETURNING user_id, balance_cents, locked_cents, updated_at
`

type UpdateWalletParams struct {
	UserID       int64 `json:"user_id"`
	BalanceCents int64 `json:"balance_cents"`
}

func (q *Queries) UpdateWallet(ctx context.Context, arg UpdateWalletParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, updateWallet, arg.UserID, arg.BalanceCents)
	var i Wallet
	err := row.Scan(
		&i.UserID,
		&i.BalanceCents,
		&i.LockedCents,
		&i.UpdatedAt,
	)
	return i, err
}
