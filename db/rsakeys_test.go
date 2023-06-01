package db

import (
	"context"
	"testing"

	"github.com/optclblast/filetagger/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateRSAkey(t *testing.T) {
	arg := CreateRSAKeyParams{
		Id:    utils.RandomID(),
		Owner: 5328,
		Key:   utils.RandomLogin(),
	}

	key, err := testQueries.CreateRSAKey(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, key)

	require.Equal(t, arg.Owner, key.Owner)
	require.Equal(t, arg.Key, key.Key)
}

func TestGetRSAKey(t *testing.T) {
	var owner int64 = 5328
	key, err := testQueries.GetRSAKey(context.Background(), owner)
	require.NoError(t, err)
	require.NotEmpty(t, key)
	require.Equal(t, "appycdlbkgtdpqk", key.Key)
}

func TestUpdateRSAKey(t *testing.T) {
	var id int64 = 5328
	arg := UpdateRSAKeyParams{
		Owner: id,
		Key:   "12345678901234567890",
	}

	key, err := testQueries.UpdateRSAKey(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, key)

	require.Equal(t, arg.Owner, key.Owner)
	require.Equal(t, arg.Key, key.Key)
}

func TestDeleteRSAKey(t *testing.T) {
	var id int64 = 1098

	err := testQueries.DeleteRSAKey(context.Background(), id)
	require.NoError(t, err)
}
