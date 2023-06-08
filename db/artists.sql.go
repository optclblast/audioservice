package db

import (
	"context"
	"fmt"
	"time"

	"github.com/optclblast/filetagger/logger"
)

type CreateArtistParams struct {
	Id   int64  `json:"id"`
	Name int64  `json:"name"`
	Bio  string `json:"bio"`
}

const createArtist = `-- name: CreateArtist :one
INSERT INTO Artists (
	id, name, bio
) VALUES (
	$1, $2, $3
) RETURNING id, name, bio
`

func (q *Queries) CreateArtist(ctx context.Context, arg CreateArtistParams) (Artist, error) {
	row := q.db.QueryRowContext(ctx, createArtist, arg.Id, arg.Name, arg.Bio)
	var i Artist
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Bio,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/file.sql.go/CreateArtist() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const getArtist = `-- name: GetArtists :one
SELECT id, name, bio FROM Artists
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetArtist(ctx context.Context, id int64, owner int64) (Artist, error) {
	row := q.db.QueryRowContext(ctx, getArtist, id, owner)
	var i Artist
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Bio,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/file.sql.go/GetArtists() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const listArtists = `--name: ListArtists :many
SELECT id, name, bio FROM Artists
WHERE name = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListArtistsParams struct {
	Limit  int32 `json:"id"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListArtists(ctx context.Context, arg ListArtistsParams, name string) ([]Artist, error) {
	rows, err := q.db.QueryContext(ctx, listArtists, arg.Limit, arg.Offset, name)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/file.sql.go/ListArtists() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	var items []Artist
	for rows.Next() {
		var i Artist
		if err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Bio,
		); err != nil {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "db/sqlc/file.sql.go/ListArtists() SCOPE",
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
			Location: "db/sqlc/file.sql.go/ListArtists() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}
	if err := rows.Err(); err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/file.sql.go/ListArtists() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	return items, nil
}

const updateArtist = `--name: UpdateArtist :one
UPDATE Artists
SET name = $2, bio = $3
WHERE id = $1
RETURNING id, name, bio 
`

type UpdateArtistParams struct {
	Id   int64  `json:"id"`
	Name int64  `json:"name"`
	Bio  string `json:"bio"`
}

func (q *Queries) UpdateArtist(ctx context.Context, arg UpdateArtistParams) (Artist, error) {
	row := q.db.QueryRowContext(ctx, updateArtist, arg.Id, arg.Name, arg.Bio)
	var i Artist
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Bio,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/file.sql.go/UpdateArtist() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const deleteArtist = `--name: DeleteArtist :exec
DELETE FROM Artists
WHERE id = $1
`

func (q *Queries) DeleteArtist(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteArtist, id)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/file.sql.go/DeleteArtist() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return err
	}
	return nil
}
