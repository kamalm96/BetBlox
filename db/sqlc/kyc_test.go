package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetKycByID(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := CreateKycParams{
		UserID:      account.ID,
		SsnLast4:    "1234",
		Dob:         time.Now(),
		Address:     "1234 Lane St",
		KycStatus:   false,
		SubmittedAt: time.Now(),
		VerifiedAt: sql.NullTime{
			Valid: false,
		},
	}

	createdKyc, err := testQueries.CreateKyc(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdKyc)
	require.Equal(t, account.ID, createdKyc.UserID)
	require.False(t, createdKyc.KycStatus)

	kycFromDB, err := testQueries.GetKyc(context.Background(), account.ID)
	require.NoError(t, err)
	require.Equal(t, createdKyc.UserID, kycFromDB.UserID)
	require.False(t, kycFromDB.KycStatus)
}

func TestUpdateKycStatus(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := CreateKycParams{
		UserID:      account.ID,
		SsnLast4:    "1234",
		Dob:         time.Now(),
		Address:     "1234 Lane St",
		KycStatus:   false,
		SubmittedAt: time.Now(),
		VerifiedAt: sql.NullTime{
			Valid: false,
		},
	}
	createdAccount, err := testQueries.CreateKyc(context.Background(), arg)
	require.NoError(t, err)
	require.False(t, createdAccount.KycStatus)

	err = testQueries.UpdateKycStatus(context.Background(), UpdateKycStatusParams{
		UserID:    account.ID,
		KycStatus: true,
	})
	require.NoError(t, err)

	updatedKycAccount, err := testQueries.GetKyc(context.Background(), account.ID)
	require.NoError(t, err)
	require.True(t, updatedKycAccount.KycStatus)
}
