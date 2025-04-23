package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateSpecificOrder(t *testing.T, userID int64, contractID int64, orderType string) Order {
	arg := CreateOrderParams{
		UserID:     userID,
		ContractID: contractID,
		OrderType:  orderType,
		OrderStyle: "limit",
		PriceCents: 1500,
		Quantity:   1,
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	return order
}

func CreateRandomTrade(t *testing.T) Trade {
	contract := CreateRandomContract(t, "YES")
	buyer := CreateRandomAccount(t)
	seller := CreateRandomAccount(t)

	buyOrder := CreateSpecificOrder(t, buyer.ID, contract.ID, "buy")
	sellOrder := CreateSpecificOrder(t, seller.ID, contract.ID, "sell")

	arg := CreateTradeParams{
		BuyOrderID:  buyOrder.ID,
		SellOrderID: sellOrder.ID,
		ContractID:  contract.ID,
		PriceCents:  1500,
		Quantity:    1,
	}

	trade, err := testQueries.CreateTrade(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trade)

	require.Equal(t, arg.BuyOrderID, trade.BuyOrderID)
	require.Equal(t, arg.SellOrderID, trade.SellOrderID)
	require.Equal(t, arg.ContractID, trade.ContractID)
	require.Equal(t, arg.PriceCents, trade.PriceCents)
	require.Equal(t, arg.Quantity, trade.Quantity)

	return trade
}

func TestCreateTrade(t *testing.T) {
	CreateRandomTrade(t)
}

func TestGetTrade(t *testing.T) {
	trade := CreateRandomTrade(t)

	result, err := testQueries.GetTrade(context.Background(), trade.ID)
	require.NoError(t, err)
	require.Equal(t, trade.ID, result.ID)
	require.Equal(t, trade.PriceCents, result.PriceCents)
	require.Equal(t, trade.BuyOrderID, result.BuyOrderID)
	require.Equal(t, trade.SellOrderID, result.SellOrderID)
}

func TestDeleteTrade(t *testing.T) {
	trade := CreateRandomTrade(t)

	err := testQueries.DeleteTrade(context.Background(), trade.ID)
	require.NoError(t, err)

	_, err = testQueries.GetTrade(context.Background(), trade.ID)
	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
}

func TestListTrades(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomTrade(t)
	}

	trades, err := testQueries.ListTrades(context.Background(), ListTradesParams{
		Limit:  5,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, trades)
	require.LessOrEqual(t, len(trades), 5)
}
