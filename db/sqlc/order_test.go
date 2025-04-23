package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomOrder(t *testing.T) Order {
	user := CreateRandomAccount(t)
	contract := CreateRandomContract(t, "YES")

	arg := CreateOrderParams{
		UserID:     user.ID,
		ContractID: contract.ID,
		OrderType:  "buy",
		OrderStyle: "limit",
		PriceCents: 1500,
		Quantity:   3,
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.UserID, order.UserID)
	require.Equal(t, arg.ContractID, order.ContractID)
	require.Equal(t, arg.OrderType, order.OrderType)
	require.Equal(t, arg.OrderStyle, order.OrderStyle)
	require.Equal(t, arg.PriceCents, order.PriceCents)
	require.Equal(t, arg.Quantity, order.Quantity)
	require.Equal(t, "open", order.Status) // assuming "open" is default

	return order
}

func TestCreateOrder(t *testing.T) {
	CreateRandomOrder(t)
}

func TestGetOrder(t *testing.T) {
	order := CreateRandomOrder(t)
	fetched, err := testQueries.GetOrder(context.Background(), order.ID)
	require.NoError(t, err)
	require.Equal(t, order.ID, fetched.ID)
	require.Equal(t, order.Status, fetched.Status)
}

func TestUpdateOrderStatus(t *testing.T) {
	order := CreateRandomOrder(t)

	err := testQueries.UpdateOrderStatus(context.Background(), UpdateOrderStatusParams{
		Status: "cancelled",
		ID:     order.ID,
	})
	require.NoError(t, err)

	updated, err := testQueries.GetOrder(context.Background(), order.ID)
	require.NoError(t, err)
	require.Equal(t, "cancelled", updated.Status)
}

func TestDeleteOrder(t *testing.T) {
	order := CreateRandomOrder(t)

	err := testQueries.DeleteOrder(context.Background(), order.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetOrder(context.Background(), order.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}
