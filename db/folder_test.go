package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/optclblast/filetagger/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateFolder(t *testing.T) {
	arg := CreateFolderParams{
		Id:          utils.RandomID(),
		Owner:       446,
		Parent:      sql.NullInt64{Valid: false},
		Name:        utils.RandomLogin(),
		AccessLevel: "DEFAULT",
		Content:     utils.RandomInt(0, 1000),
		Tag:         utils.RandomLogin(),
	}

	account, err := testQueries.CreateFolder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Id, account.Id)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.AccessLevel, account.AccessLevel)
	require.Equal(t, arg.Content, account.Content)
	require.Equal(t, arg.Tag, account.Tag)

	require.NotZero(t, account.Id)
	require.NotZero(t, account.CreatedAt)
}
