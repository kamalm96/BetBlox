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
		Dob:         time.Time{},
		Address:     "1234 Lane St",
		KycStatus:   false,
		SubmittedAt: time.Time{},
		VerifiedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}

	accountKys, err := testQueries.CreateKyc(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountKys)
	require.Equal(t, accountKys.UserID, account.ID)

}
