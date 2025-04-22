package db

import (
	"context"
	"github.com/kamalm96/backend/db/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

func HashPassword(n string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(n), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password")
	}
	return hashedPassword
}

func CreateRandomAccount(t *testing.T) CreateAccountRow {
	arg := CreateAccountParams{
		Email:        utils.RandomEmail(),
		Username:     utils.RandomString(6),
		PasswordHash: string(HashPassword("password")),
	}
	userRow, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userRow)
	require.Equal(t, arg.Username, userRow.Username)
	require.Equal(t, arg.Email, userRow.Email)
	require.NotZero(t, userRow.ID)
	require.NotZero(t, userRow.CreatedAt)

	return userRow
}

func TestCreateAccountRow(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetUser(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Username, account2.Username)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomAccount(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	err := testQueries.DeleteUser(context.Background(), account.ID)
	require.NoError(t, err)

	deletedAccount, err := testQueries.GetUser(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, deletedAccount)
}

func TestUpdateUserPassword(t *testing.T) {
	account := CreateRandomAccount(t)
	require.NotEmpty(t, account)

	hashBytes := HashPassword("12345")
	newPasswordHash := string(hashBytes)

	err := testQueries.UpdateUserPassword(context.Background(), UpdateUserPasswordParams{
		PasswordHash: newPasswordHash,
		Email:        account.Email,
	})
	require.NoError(t, err)

	updatedAccount, err := testQueries.GetUser(context.Background(), account.ID)
	require.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(updatedAccount.PasswordHash), []byte("12345"))
	require.NoError(t, err)
}
