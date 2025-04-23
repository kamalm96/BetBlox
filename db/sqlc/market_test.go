package db

import (
	"context"
	"github.com/kamalm96/backend/db/utils"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func CreateRandomMarket(t *testing.T) Market {
	statuses := []string{"open", "closed", "resolved"}

	arg := CreateMarketParams{
		Title:       utils.RandomString(10),
		Description: utils.RandomString(50),
		Category:    utils.RandomString(5),
		Status:      statuses[rand.Intn(len(statuses))],
		ClosesAt:    time.Now(),
		ResolvesAt:  time.Now(),
	}

	market, err := testQueries.CreateMarket(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, market)

	return market
}

func CreateOpenMarket(t *testing.T) Market {
	arg := CreateMarketParams{
		Title:       utils.RandomString(10),
		Description: utils.RandomString(50),
		Category:    utils.RandomString(5),
		Status:      "open",
		ClosesAt:    time.Now(),
		ResolvesAt:  time.Now(),
	}

	market, err := testQueries.CreateMarket(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, market)

	return market
}

func TestCreateRandomMarket(t *testing.T) {
	CreateRandomMarket(t)
}

func TestGetMarket(t *testing.T) {
	market := CreateRandomMarket(t)
	market2, err := testQueries.GetMarket(context.Background(), market.ID)
	require.NoError(t, err)
	require.NotEmpty(t, market2)
	require.Equal(t, market2.ID, market.ID)
	require.Equal(t, market2.Category, market.Category)
	require.Equal(t, market2.CreatedAt, market.CreatedAt)
	require.Equal(t, market2.Title, market.Title)
	require.Equal(t, market2.Status, market.Status)
	require.Equal(t, market2.Description, market.Description)
}

func TestListAllMarkets(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomMarket(t)
	}

	arg := ListAllMarketsParams{
		Limit:  5,
		Offset: 2,
	}
	markets, err := testQueries.ListAllMarkets(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, markets, int(arg.Limit))
}

func TestListAllOpenMarkets(t *testing.T) {
	for i := 0; i < 20; i++ {
		CreateOpenMarket(t)
	}

	arg := ListOpenMarketsParams{
		Limit:  20,
		Offset: 2,
	}
	markets, err := testQueries.ListOpenMarkets(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, markets)

	for _, m := range markets {
		require.Equal(t, "open", strings.ToLower(m.Status))
	}
}

func TestUpdateResolution(t *testing.T) {
	market := CreateRandomMarket(t)

	arg := ResolveMarketParams{
		Status: "resolved",
		ID:     market.ID,
	}

	err := testQueries.ResolveMarket(context.Background(), arg)
	require.NoError(t, err)
	resolvedMarket, err := testQueries.GetMarket(context.Background(), arg.ID)
	require.NoError(t, err)
	require.Equal(t, market.ID, resolvedMarket.ID)
	require.Equal(t, resolvedMarket.Status, "resolved")
}
