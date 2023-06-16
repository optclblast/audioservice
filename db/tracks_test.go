package db

import (
	"context"
	"testing"

	"github.com/optclblast/audioservice/utils"
	"github.com/stretchr/testify/require"
)

func createRandomRSAKey(t *testing.T) RSAKey {
	params := ListAccountsParams{
		Limit:  50,
		Offset: 0,
	}
	accs, err := testQueries.ListAccounts(context.Background(), params)
	require.NoError(t, err)
	owner := accs[utils.RandomNumber(0, int64(len(accs)-1))]

	arg := CreateRSAKeyParams{
		Owner: owner.Id,
		Key:   utils.RandomLogin(),
	}

	key, err := testQueries.CreateRSAKey(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, key)

	require.Equal(t, arg.Owner, key.Owner)
	require.Equal(t, arg.Key, key.Key)
	return key
}
func TestCreateRSAkey(t *testing.T) {
	createRandomRSAKey(t)
}

func TestGetRSAKey(t *testing.T) {
	rkey := createRandomRSAKey(t)
	key, err := testQueries.GetRSAKey(context.Background(), rkey.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, key)
	require.Equal(t, rkey.Key, key.Key)
}

func TestUpdateRSAKey(t *testing.T) {
	params := ListAccountsParams{
		Limit:  50,
		Offset: 0,
	}
	accs, err := testQueries.ListAccounts(context.Background(), params)
	require.NoError(t, err)
	owner := accs[utils.RandomNumber(0, int64(len(accs)-1))]

	arg := UpdateRSAKeyParams{
		Owner: owner.Id,
		Key:   utils.RandomLogin(),
	}

	key, err := testQueries.UpdateRSAKey(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, key)

	require.Equal(t, arg.Owner, key.Owner)
	require.Equal(t, arg.Key, key.Key)
}

func TestDeleteRSAKey(t *testing.T) {
	paramsA := ListAccountsParams{
		Limit:  50,
		Offset: 0,
	}
	accs, _ := testQueries.ListAccounts(context.Background(), paramsA)
	owner := accs[utils.RandomNumber(0, int64(len(accs)-1))]

	err := testQueries.DeleteRSAKey(context.Background(), owner.Id)
	require.NoError(t, err)
}
