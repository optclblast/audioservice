package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/optclblast/filetagger/utils"
	"github.com/stretchr/testify/require"
)

func createRandomFolder(t *testing.T) Folder {
	params := ListAccountsParams{
		Limit:  50,
		Offset: 0,
	}
	accs, _ := testQueries.ListAccounts(context.Background(), params)
	owner := accs[utils.RandomNumber(0, int64(len(accs)-1))]

	arg := CreateFolderParams{
		Owner:  owner.Id,
		Parent: sql.NullInt64{Valid: false},
		Name:   utils.RandomLogin(),
		Path:   utils.RandomLogin(),
		Tag:    utils.RandomLogin(),
	}

	folder, err := testQueries.CreateFolder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, folder)

	require.Equal(t, arg.Owner, folder.Owner)
	require.Equal(t, arg.Name, folder.Name)
	require.Equal(t, arg.Path, folder.Path)
	require.Equal(t, arg.Tag, folder.Tag)

	require.NotZero(t, folder.Id)
	require.NotZero(t, folder.CreatedAt)
	return folder
}

func TestCreateFolder(t *testing.T) {
	createRandomFolder(t)
}

func TestGetFolder(t *testing.T) {
	randFolder := createRandomFolder(t)

	folder, err := testQueries.GetFolder(context.Background(), randFolder.Id, randFolder.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, folder)
	require.Equal(t, randFolder.Id, folder.Id)
	require.Equal(t, randFolder.Owner, folder.Owner)
	require.NotZero(t, folder.Id)
	require.NotZero(t, folder.CreatedAt)
}

func TestListFolder(t *testing.T) {
	randFolder := createRandomFolder(t)

	arg := ListFoldersParams{
		Limit:  int32(utils.RandomInt(1, 3)),
		Offset: int32(utils.RandomInt(1, 2)),
	}

	folders, err := testQueries.ListFolders(context.Background(), arg, randFolder.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, folders)
}

func TestUpdateFolder(t *testing.T) {
	randFolder := createRandomFolder(t)

	arg := UpdateFolderParams{
		Id:   randFolder.Id,
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
	paramsA := ListAccountsParams{
		Limit:  50,
		Offset: 0,
	}
	accs, _ := testQueries.ListAccounts(context.Background(), paramsA)
	owner := accs[utils.RandomNumber(0, int64(len(accs)-1))]

	params := ListFoldersParams{
		Limit:  50,
		Offset: 0,
	}
	var folders []Folder
	for {
		var err error
		folders, err = testQueries.ListFolders(context.Background(), params, owner.Id)
		if err == nil && len(folders) > 0 {
			break
		}
	}
	folder := accs[utils.RandomNumber(0, int64(len(folders)-1))]

	err := testQueries.DeleteFolder(context.Background(), folder.Id)
	require.NoError(t, err)
}
