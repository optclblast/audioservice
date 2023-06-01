package db

import (
	"context"
	"testing"

	"github.com/optclblast/filetagger/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateFile(t *testing.T) {
	arg := CreateFileParams{
		Id:     utils.RandomID(),
		Owner:  6845,
		Parent: 3373,
		Name:   utils.RandomLogin(),
		Path:   utils.RandomLogin(),
		Tag:    utils.RandomLogin(),
	}

	file, err := testQueries.CreateFile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, file)

	require.Equal(t, arg.Id, file.Id)
	require.Equal(t, arg.Owner, file.Owner)
	require.Equal(t, arg.Name, file.Name)
	require.Equal(t, arg.Path, file.Path)
	require.Equal(t, arg.Tag, file.Tag)

	require.NotZero(t, file.Id)
	require.NotZero(t, file.CreatedAt)
}

func TestGetFile(t *testing.T) {
	var id int64 = 584
	var owner int64 = 6845
	folder, err := testQueries.GetFile(context.Background(), id, owner)
	require.NoError(t, err)
	require.NotEmpty(t, folder)
	require.Equal(t, id, folder.Id)
	require.Equal(t, owner, folder.Owner)
	require.NotZero(t, folder.Id)
}

func TestListFile(t *testing.T) {
	var owner int64 = 711

	arg := ListFilesParams{
		Limit:  int32(utils.RandomInt(1, 5)),
		Offset: int32(utils.RandomInt(1, 2)),
	}

	folders, err := testQueries.ListFiles(context.Background(), arg, owner)
	require.NoError(t, err)
	require.NotEmpty(t, folders)
}

func TestUpdateFile(t *testing.T) {
	var id int64 = 4700
	arg := UpdateFileParams{
		Id:   id,
		Name: "PAPICH",
		Path: utils.RandomLogin(),
		Tag:  utils.RandomLogin(),
	}

	folder, err := testQueries.UpdateFile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, folder)

	require.Equal(t, arg.Id, folder.Id)
	require.Equal(t, arg.Name, folder.Name)
	require.Equal(t, arg.Path, folder.Path)
	require.Equal(t, arg.Tag, folder.Tag)

	require.NotZero(t, folder.Id)
}

func TestDeleteFile(t *testing.T) {
	var id int64 = 5103

	err := testQueries.DeleteFile(context.Background(), id)
	require.NoError(t, err)
}
