package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Id:       144,
		Login:    "qwer67",
		Password: "qweq98rrte",
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
