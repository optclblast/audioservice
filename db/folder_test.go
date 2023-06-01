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
		Id:     utils.RandomID(),
		Owner:  6845,
		Parent: sql.NullInt64{Valid: false},
		Name:   utils.RandomLogin(),
		Path:   utils.RandomLogin(),
		Tag:    utils.RandomLogin(),
	}

	folder, err := testQueries.CreateFolder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, folder)

	require.Equal(t, arg.Id, folder.Id)
	require.Equal(t, arg.Owner, folder.Owner)
	require.Equal(t, arg.Name, folder.Name)
	require.Equal(t, arg.Path, folder.Path)
	require.Equal(t, arg.Tag, folder.Tag)

	require.NotZero(t, folder.Id)
	require.NotZero(t, folder.CreatedAt)
}

func TestGetFolder(t *testing.T) {
	var id int64 = 3373
	var owner int64 = 6845
	folder, err := testQueries.GetFolder(context.Background(), id, owner)
	require.NoError(t, err)
	require.NotEmpty(t, folder)
	require.Equal(t, id, folder.Id)
	require.Equal(t, owner, folder.Owner)
	require.NotZero(t, folder.Id)
	require.NotZero(t, folder.CreatedAt)
}

func TestListFolder(t *testing.T) {
	var owner int64 = 6845

	arg := ListFoldersParams{
		Limit:  int32(utils.RandomInt(1, 3)),
		Offset: int32(utils.RandomInt(1, 2)),
	}

	folders, err := testQueries.ListFolders(context.Background(), arg, owner)
	require.NoError(t, err)
	require.NotEmpty(t, folders)
}

func TestUpdateFolder(t *testing.T) {
	var id int64 = 606
	arg := UpdateFolderParams{
		Id:   id,
		Name: "PAPICH",
		Path: utils.RandomLogin(),
		Tag:  utils.RandomLogin(),
	}

	folder, err := testQueries.UpdateFolder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, folder)

	require.Equal(t, arg.Id, folder.Id)
	require.Equal(t, arg.Name, folder.Name)
	require.Equal(t, arg.Path, folder.Path)
	require.Equal(t, arg.Tag, folder.Tag)

	require.NotZero(t, folder.Id)
}

func TestDeleteFolder(t *testing.T) {
	var id int64 = 418

	err := testQueries.DeleteFolder(context.Background(), id)
	require.NoError(t, err)
}
