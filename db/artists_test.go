package db

import (
	"context"
	"testing"

	"github.com/optclblast/filetagger/utils"
	"github.com/stretchr/testify/require"
)

func createRandomFile(t *testing.T) File {
	params := ListAccountsParams{
		Limit:  50,
		Offset: 0,
	}
	accs, err := testQueries.ListAccounts(context.Background(), params)
	require.NoError(t, err)
	owner := accs[utils.RandomNumber(0, int64(len(accs)-1))]

	dirParams := ListFoldersParams{
		Limit:  50,
		Offset: 0,
	}
	dirs, err := testQueries.ListFolders(context.Background(), dirParams, owner.Id)
	require.NotEmpty(t, dirs)
	require.NoError(t, err)
	dir := dirs[utils.RandomNumber(0, int64(len(dirs)-1))]
	require.Equal(t, dir.Owner, owner.Id)

	arg := CreateFileParams{
		Owner:  owner.Id,
		Parent: dir.Id,
		Name:   utils.RandomLogin(),
		Path:   utils.RandomLogin(),
		Tag:    utils.RandomLogin(),
	}

	file, err := testQueries.CreateFile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, file)

	require.Equal(t, arg.Owner, file.Owner)
	require.Equal(t, arg.Name, file.Name)
	require.Equal(t, arg.Path, file.Path)
	require.Equal(t, arg.Tag, file.Tag)

	require.NotZero(t, file.Id)
	require.NotZero(t, file.CreatedAt)

	return file
}
func TestCreateFile(t *testing.T) {
	createRandomFile(t)
}

func TestGetFile(t *testing.T) {
	randFile := createRandomFile(t)
	folder, err := testQueries.GetFile(context.Background(), randFile.Id, randFile.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, folder)
	require.Equal(t, randFile.Id, folder.Id)
	require.Equal(t, randFile.Owner, folder.Owner)
	require.NotZero(t, folder.Id)
}

func TestListFile(t *testing.T) {
	randFile := createRandomFile(t)

	arg := ListFilesParams{
		Limit:  int32(utils.RandomInt(1, 5)),
		Offset: 0,
	}

	files, err := testQueries.ListFiles(context.Background(), arg, randFile.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, files)
}

func TestUpdateFile(t *testing.T) {
	randFile := createRandomFile(t)

	arg := UpdateFileParams{
		Id:   randFile.Id,
		Name: "random power",
		Path: utils.RandomLogin(),
		Tag:  utils.RandomLogin(),
	}

	file, err := testQueries.UpdateFile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, file)

	require.Equal(t, arg.Id, file.Id)
	require.Equal(t, arg.Name, file.Name)
	require.Equal(t, arg.Path, file.Path)
	require.Equal(t, arg.Tag, file.Tag)

	require.NotZero(t, file.Id)
}

func TestDeleteFile(t *testing.T) {
	paramsA := ListAccountsParams{
		Limit:  50,
		Offset: 0,
	}
	accs, _ := testQueries.ListAccounts(context.Background(), paramsA)
	owner := accs[utils.RandomNumber(0, int64(len(accs)-1))]

	params := ListFilesParams{
		Limit:  50,
		Offset: 0,
	}
	var files []File
	for {
		var err error
		files, err = testQueries.ListFiles(context.Background(), params, owner.Id)
		if err == nil && len(files) > 0 {
			break
		}
	}
	file := accs[utils.RandomNumber(0, int64(len(files)-1))]

	err := testQueries.DeleteFile(context.Background(), file.Id)
	require.NoError(t, err)
}
