package db

import (
	"context"
	"database/sql"
	"github.com/kamalm96/backend/db/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := CreateWalletParams{
		UserID:       account.ID,
		BalanceCents: int64(utils.RandomInt(50, 200)),
		LockedCents:  int64(utils.RandomInt(50, 200)),
		UpdatedAt: sql.NullTime{
			Valid: false,
		},
	}

	createdWallet, err := testQueries.CreateWallet(context.Background(), arg)
	gettingWallet, err := testQueries.GetWallet(context.Background(), createdWallet.UserID)
	require.NoError(t, err)
	require.Equal(t, gettingWallet.UserID, account.ID)
	require.Equal(t, gettingWallet.BalanceCents, createdWallet.BalanceCents)
	require.Equal(t, gettingWallet.LockedCents, createdWallet.LockedCents)

}

func TestUpdateRealWallet(t *testing.T) {
	account := CreateRandomAccount(t)
	arg := CreateWalletParams{
		UserID:       account.ID,
		BalanceCents: int64(utils.RandomInt(50, 200)),
		LockedCents:  int64(utils.RandomInt(50, 200)),
		UpdatedAt: sql.NullTime{
			Valid: false,
		},
	}

	createdWallet, err := testQueries.CreateWallet(context.Background(), arg)
	require.NoError(t, err)

	balanceAdded := 22

	updatedBalancePlus, err := testQueries.UpdateWallet(context.Background(), UpdateWalletParams{
		UserID:       createdWallet.UserID,
		BalanceCents: createdWallet.BalanceCents + int64(balanceAdded),
	})
	require.NoError(t, err)
	require.Equal(t, createdWallet.BalanceCents+int64(balanceAdded), updatedBalancePlus.BalanceCents)
	balanceRemoved := int64(22)
	updatedBalanceMinus, err := testQueries.UpdateWallet(context.Background(), UpdateWalletParams{
		UserID:       account.ID,
		BalanceCents: createdWallet.BalanceCents - balanceRemoved,
	})

	require.NoError(t, err)
	require.Equal(t, updatedBalanceMinus.UserID, account.ID)
	require.Equal(t, createdWallet.BalanceCents-balanceRemoved, updatedBalanceMinus.BalanceCents)
}

func TestUpdateLockedBalance(t *testing.T) {

	account := CreateRandomAccount(t)
	arg := CreateWalletParams{
		UserID:       account.ID,
		BalanceCents: int64(utils.RandomInt(50, 200)),
		LockedCents:  int64(utils.RandomInt(50, 200)),
		UpdatedAt: sql.NullTime{
			Valid: false,
		},
	}

	createdWallet, err := testQueries.CreateWallet(context.Background(), arg)
	require.NoError(t, err)
	lockedPlus := int64(utils.RandomInt(1, 100))
	updatedLockedBalance, err := testQueries.UpdateLocked(context.Background(), UpdateLockedParams{
		UserID:      createdWallet.UserID,
		LockedCents: createdWallet.LockedCents + lockedPlus,
	})
	require.NoError(t, err)
	require.Equal(t, createdWallet.UserID, account.ID)
	require.Equal(t, createdWallet.LockedCents+lockedPlus, updatedLockedBalance.LockedCents)

}
