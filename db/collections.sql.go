package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/optclblast/filetagger/logger"
)

type CreateCollectionParams struct {
	Name        string        `json:"name"`
	Author      int64         `json:"author"`
	FtAuthors   sql.NullInt64 `json:"ft_authors"`
	Type        string        `json:"type"`
	Discription string        `json:"discription"`
	Lenght      string        `json:"lenght"`
	Label       string        `json:"label"`
	Date        time.Time     `json:"date"`
}

const createCollection = `-- name: CreateCollection :one
INSERT INTO Collections (
	name, author, ft_authors, type, discription, lenght, label, date
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, name, author, ft_authors, type, discription, lenght, label, date
`

func (q *Queries) CreateCollection(ctx context.Context, arg CreateCollectionParams) (Collection, error) {
	row := q.db.QueryRowContext(ctx, createCollection, arg.Name, arg.Author, arg.FtAuthors, arg.Type, arg.Discription, arg.Lenght, arg.Label, arg.Date)

	var i Collection
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Author,
		&i.FtAuthors,
		&i.Type,
		&i.Discription,
		&i.Lenght,
		&i.Label,
		&i.Date,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/Collection.sql.go/CreateCollection() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const getCollection = `-- name: GetCollection :one
SELECT id, name, author, ft_authors, type, discription, lenght, label, date FROM Collections
WHERE id = $1 AND author = $2 LIMIT 1
`

func (q *Queries) GetCollection(ctx context.Context, id int64, author int64) (Collection, error) {
	row := q.db.QueryRowContext(ctx, getCollection, id, author)
	var i Collection
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Author,
		&i.FtAuthors,
		&i.Type,
		&i.Discription,
		&i.Lenght,
		&i.Label,
		&i.Date,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/Collection.sql.go/GetCollection() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const listCollections = `--name: ListCollections :many
SELECT id, name, author, ft_authors, type, discription, lenght, label, date FROM Collections
WHERE author = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCollectionsParams struct {
	Limit  int32 `json:"id"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCollections(ctx context.Context, arg ListCollectionsParams, author int64) ([]Collection, error) {
	rows, err := q.db.QueryContext(ctx, listCollections, arg.Limit, arg.Offset, author)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/Collection.sql.go/ListCollections() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Author,
			&i.FtAuthors,
			&i.Type,
			&i.Discription,
			&i.Lenght,
			&i.Label,
			&i.Date,
		); err != nil {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "db/sqlc/Collection.sql.go/ListCollections() SCOPE",
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
			Location: "db/sqlc/Collection.sql.go/ListCollections() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}
	if err := rows.Err(); err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/Collection.sql.go/ListCollections() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	return items, nil
}

const updateCollection = `--name: UpdateCollection :one
UPDATE Collections
SET name = $2, author = $3, ft_authors = $4, type = $5, discription = $6, lenght = $7, label = $8, date = $9
WHERE id = $1
RETURNING name, author, ft_authors, type, discription, lenght, label, date
`

type UpdateCollectionParams struct {
	Name        string        `json:"name"`
	Author      int64         `json:"author"`
	FtAuthors   sql.NullInt64 `json:"ft_authors"`
	Type        string        `json:"type"`
	Discription string        `json:"discription"`
	Lenght      string        `json:"lenght"`
	Label       string        `json:"label"`
	Date        time.Time     `json:"date"`
}

func (q *Queries) UpdateCollection(ctx context.Context, arg UpdateCollectionParams) (Collection, error) {
	row := q.db.QueryRowContext(ctx, updateCollection, arg.Name, arg.Author, arg.FtAuthors, arg.Type, arg.Discription, arg.Lenght, arg.Label, arg.Date)
	var i Collection
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Author,
		&i.FtAuthors,
		&i.Type,
		&i.Discription,
		&i.Lenght,
		&i.Label,
		&i.Date,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/Collection.sql.go/UpdateCollection() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const deleteCollection = `--name: DeleteCollection :exec
DELETE FROM Collection
WHERE id = $1
`

func (q *Queries) DeleteCollection(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCollection, id)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/Collection.sql.go/DeleteCollection() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return err
	}
	return nil
}
