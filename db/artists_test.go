package db

import (
	"context"
	"testing"

	"github.com/optclblast/audioservice/utils"
	"github.com/stretchr/testify/require"
)

func createRandomArtist(t *testing.T) Artist {
	arg := CreateArtistParams{
		Name: utils.RandomLogin(),
		Bio:  "some bio",
	}

	artist, err := testQueries.CreateArtist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, artist)

	require.Equal(t, arg.Bio, artist.Bio)
	require.Equal(t, arg.Name, artist.Name)

	require.NotZero(t, artist.Id)

	return artist
}
func TestCreateArtist(t *testing.T) {
	createRandomArtist(t)
}

func TestGetArtist(t *testing.T) {
	randArtist := createRandomArtist(t)
	artist, err := testQueries.GetArtist(context.Background(), randArtist.Id)
	require.NoError(t, err)
	require.NotEmpty(t, artist)
	require.Equal(t, randArtist.Id, artist.Id)
	require.Equal(t, randArtist.Name, artist.Name)
	require.Equal(t, randArtist.Bio, artist.Bio)
	require.NotZero(t, artist.Id)
}

func TestListArtists(t *testing.T) {
	randArtist := createRandomArtist(t)

	arg := ListArtistsParams{
		Name:   randArtist.Name,
		Limit:  int32(utils.RandomInt(1, 5)),
		Offset: 0,
	}

	Artists, err := testQueries.ListArtists(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Artists)
}

func TestUpdateArtist(t *testing.T) {
	randArtist := createRandomArtist(t)

	arg := UpdateArtistParams{
		Id:   randArtist.Id,
		Name: "NEW_NAME",
		Bio:  utils.RandomLogin(),
	}

	Artist, err := testQueries.UpdateArtist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Artist)

	require.Equal(t, arg.Id, Artist.Id)
	require.Equal(t, arg.Name, Artist.Name)
	require.Equal(t, arg.Bio, Artist.Bio)

	require.NotZero(t, Artist.Id)
}

func TestDeleteArtist(t *testing.T) {
	randArtist := createRandomArtist(t)
	artist, err := testQueries.GetArtist(context.Background(), randArtist.Id)

	err = testQueries.DeleteArtist(context.Background(), artist.Id)
	require.NoError(t, err)
}
