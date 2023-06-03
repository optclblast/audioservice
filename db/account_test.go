package db

import (
	"context"
	"testing"

	"github.com/optclblast/filetagger/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Login:    utils.RandomLogin(),
		Password: utils.RandomPassword(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Login, account.Login)
	require.Equal(t, arg.Password, account.Password)

	require.NotZero(t, account.Id)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdAcc := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), createdAcc.Id)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createdAcc.Id, account.Id)
	require.Equal(t, createdAcc.Login, account.Login)
	require.Equal(t, createdAcc.Password, account.Password)
}

func TestListAccounts(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  int32(utils.RandomInt(1, 3)),
		Offset: int32(utils.RandomInt(1, 2)),
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}

func TestDeleteAccount(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 67587907098568709)
	require.NoError(t, err)
}
