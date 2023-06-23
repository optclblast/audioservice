package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/optclblast/audioservice/logger"
)

type CreateTrackParams struct {
	Name      string        `json:"name"`
	Artist    int64         `json:"artist"`
	FtArtists sql.NullInt64 `json:"ft_artists"`
	Album     int64         `json:"album"`
	Location  string        `json:"location"`
}

const createTrack = `-- name: CreateTrack :one
INSERT INTO Tracks (
	name, artist, ft_artists, album, location
) VALUES (
	$1, $2, $3, $4, $5, $6
) RETURNING id, name, artist, ft_artists, album, location
`

func (q *Queries) CreateTrack(ctx context.Context, arg CreateTrackParams) (Track, error) {
	row := q.db.QueryRowContext(ctx, createTrack, arg.Name, arg.Artist, arg.Artist)
	var i Track
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Artist,
		&i.FtArtists,
		&i.Album,
		&i.Location,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/Tracks.sql.go/CreateTrack() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const getTrack = `-- name: GetTrack :one
SELECT key FROM Tracks
WHERE owner = $1 
ORDER BY id
LIMIT 1
`

func (q *Queries) GetTrack(ctx context.Context, owner int64) (Track, error) {
	row := q.db.QueryRowContext(ctx, getTrack, owner)
	var i Track
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Artist,
		&i.FtArtists,
		&i.Album,
		&i.Location,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/Tracks.sql.go/GetTrack() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const listTracks = `--name: ListCollections :many
SELECT id, name, author, ft_authors, type, discription, lenght, label, date FROM Collections
WHERE author = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTracksParams struct {
	Limit  int32 `json:"id"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTracks(ctx context.Context, arg ListTracksParams, author int64) ([]Track, error) {
	rows, err := q.db.QueryContext(ctx, listTracks, arg.Limit, arg.Offset, author)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/Tracks.sql.go/ListTracks() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	var items []Track
	for rows.Next() {
		var i Track
		if err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Artist,
			&i.FtArtists,
			&i.Album,
			&i.Location,
		); err != nil {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "db/sqlc/Tracks.sql.go/ListCollections() SCOPE",
				Content:  fmt.Sprint(err),
			})
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/Tracks.sql.go/ListCollections() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}
	if err := rows.Err(); err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/Tracks.sql.go/ListCollections() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	return items, nil
}

const updateTrack = `--name: UpdateTrack :one
UPDATE Tracks
SET key = $2
WHERE owner = $1
RETURNING owner, key
`

type UpdateTrackParams struct {
	Owner int64  `json:"owner"`
	Key   string `json:"key"`
}

func (q *Queries) UpdateTrack(ctx context.Context, arg UpdateTrackParams) (Track, error) {
	row := q.db.QueryRowContext(ctx, updateTrack, arg.Owner, arg.Key)
	var i Track
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Artist,
		&i.FtArtists,
		&i.Album,
		&i.Location,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/Tracks.sql.go/UpdateTrack() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const deleteTrack = `--name: DeleteTrack :exec
DELETE FROM Tracks
WHERE owner = $1
`

func (q *Queries) DeleteTrack(ctx context.Context, owner int64) error {
	_, err := q.db.ExecContext(ctx, deleteTrack, owner)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/Tracks.sql.go/DeleteTrack() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return err
	}
	return nil
}
