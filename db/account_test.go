package db

import (
	"context"
	"testing"

	"github.com/optclblast/filetagger/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Id:       utils.RandomID(),
		Login:    utils.RandomLogin(),
		Password: utils.RandomPassword(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Id, account.Id)
	require.Equal(t, arg.Login, account.Login)
	require.Equal(t, arg.Password, account.Password)

	require.NotZero(t, account.Id)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	var id int64 = 144
	account, err := testQueries.GetAccount(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, id, account.Id)
	require.Equal(t, "qwer67", account.Login)
	require.Equal(t, "qweq98rrte", account.Password)
}

func TestListAccounts(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  int32(utils.RandomInt(1, 10)),
		Offset: int32(utils.RandomInt(1, 10)),
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}

func TestDeleteAccount(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 67587907098568709)
	require.NoError(t, err)
}
