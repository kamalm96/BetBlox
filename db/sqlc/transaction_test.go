package db

import (
	"context"
	"github.com/kamalm96/backend/db/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomTransaction(t *testing.T) Transaction {
	account := CreateRandomAccount(t)
	arg := CreateTransactionParams{
		ID:           int64(utils.RandomInt(1, 100000)),
		UserID:       account.ID,
		Type:         "buy",
		AmountCents:  int64(utils.RandomInt(100, 1000)),
		BalanceAfter: int64(utils.RandomInt(100, 1000)),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.Equal(t, transaction.UserID, account.ID)

	return transaction

}
func TestCreateRandomTransaction(t *testing.T) {
	CreateRandomTransaction(t)
}

func TestGetTransaction(t *testing.T) {
	tx := CreateRandomTransaction(t)

	transaction, err := testQueries.GetTransactionById(context.Background(), tx.ID)
	require.NoError(t, err)
	require.Equal(t, transaction.UserID, tx.UserID)
	require.Equal(t, transaction.ID, tx.ID)
	require.Equal(t, transaction.Type, tx.Type)

}

func TestListAllTransactions(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomTransaction(t)
	}

	arg := GetAllTransactionsParams{
		Limit:  5,
		Offset: 1,
	}

	transactions, err := testQueries.GetAllTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transactions, 5)

}
