package db

import (
	"context"
	"github.com/kamalm96/backend/db/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomContract(t *testing.T, n string) Contract {
	arg := CreateContractParams{
		ContractType: n,
		PriceCents:   int32(utils.RandomInt(20, 100)),
	}

	contract, err := testQueries.CreateContract(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contract)
	require.Equal(t, contract.ContractType, arg.ContractType)
	require.Equal(t, contract.PriceCents, arg.PriceCents)

	return contract

}

func TestCreateRandomContract(t *testing.T) {
	CreateRandomContract(t, "YES")
	CreateRandomContract(t, "NO")
}

func TestGetContract(t *testing.T) {
	contract := CreateRandomContract(t, "YES")

	gettingContract, err := testQueries.GetContract(context.Background(), contract.ID)
	require.NoError(t, err)
	require.Equal(t, gettingContract.ID, contract.ID)
	require.Equal(t, gettingContract.ContractType, contract.ContractType)
	require.Equal(t, gettingContract.PriceCents, contract.PriceCents)
	require.Equal(t, gettingContract.MarketID, contract.MarketID)
	require.Equal(t, gettingContract.Volume, contract.Volume)

}

func TestListContracts(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomContract(t, "YES")
	}
	arg := ListContractsParams{
		Limit:  5,
		Offset: 1,
	}

	contracts, err := testQueries.ListContracts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, contracts, int(arg.Limit))
}

func TestDeleteContract(t *testing.T) {
	contract := CreateRandomContract(t, "NO")
	require.NotEmpty(t, contract)

	err := testQueries.DeleteContract(context.Background(), contract.ID)
	require.NoError(t, err)

	deletedContract, err := testQueries.GetContract(context.Background(), contract.ID)
	require.Error(t, err)
	require.Empty(t, deletedContract)
}
