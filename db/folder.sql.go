package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/optclblast/filetagger/logger"
)

// operations with folders
type CreateFolderParams struct {
	Id     int64         `json:"id"`
	Owner  int64         `json:"owner"`
	Parent sql.NullInt64 `json:"parent"`
	Name   string        `json:"name"`
	Path   string        `json:"path"`
	Tag    string        `json:"tag"`
}

const createFolder = `-- name: CreateFolder :one
INSERT INTO folder (
	id,
	owner,
	parent,
	name,
	created_at,
	path,
	tag
) VALUES (
	$1, $2, $3, $4, $5, $6, $7
) RETURNING id, owner, parent, name, created_at, path, tag
`

func (q *Queries) CreateFolder(ctx context.Context, arg CreateFolderParams) (Folder, error) {
	row := q.db.QueryRowContext(ctx, createFolder, arg.Id, arg.Owner, arg.Parent, arg.Name, time.Now(), arg.Path, arg.Tag)

	var i Folder
	err := row.Scan(
		&i.Id,
		&i.Owner,
		&i.Parent,
		&i.Name,
		&i.CreatedAt,
		&i.Path,
		&i.Tag,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/folder.sql.go/CreateFolder() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const getFolder = `-- name: GetFolder :one
SELECT id, owner, parent, name, created_at, path, tag FROM folder
WHERE id = $1 AND owner = $2 LIMIT 1
`

func (q *Queries) GetFolder(ctx context.Context, id int64, owner int64) (Folder, error) {
	row := q.db.QueryRowContext(ctx, getFolder, id, owner)
	var i Folder
	err := row.Scan(
		&i.Id,
		&i.Owner,
		&i.Parent,
		&i.Name,
		&i.CreatedAt,
		&i.Path,
		&i.Tag,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/folder.sql.go/GetFolder() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const listFolders = `--name: ListFolders :many
SELECT id, owner, parent, name, created_at, path, tag FROM folder
WHERE owner = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListFoldersParams struct {
	Limit  int32 `json:"id"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFolders(ctx context.Context, arg ListFoldersParams, owner int64) ([]Folder, error) {
	rows, err := q.db.QueryContext(ctx, listFolders, arg.Limit, arg.Offset, owner)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/folder.sql.go/ListFolders() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	var items []Folder
	for rows.Next() {
		var i Folder
		if err := rows.Scan(
			&i.Id,
			&i.Owner,
			&i.Parent,
			&i.Name,
			&i.CreatedAt,
			&i.Path,
			&i.Tag,
		); err != nil {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "db/sqlc/folder.sql.go/ListFolders() SCOPE",
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
			Location: "db/sqlc/folder.sql.go/ListFolders() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}
	if err := rows.Err(); err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/folder.sql.go/ListFolders() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	return items, nil
}

const updateFolder = `--name: UpdateFolder :one
UPDATE folder
SET name = $2, path = $3, tag = $4
WHERE id = $1
RETURNING id, name, path, tag
`

type UpdateFolderParams struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Tag  string `json:"tag"`
}

func (q *Queries) UpdateFolder(ctx context.Context, arg UpdateFolderParams) (Folder, error) {
	row := q.db.QueryRowContext(ctx, updateFolder, arg.Id, arg.Name, arg.Path, arg.Tag)
	var i Folder
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Path,
		&i.Tag,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/folder.sql.go/UpdateFolders() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const deleteFolder = `--name: DeleteFolder :exec
DELETE FROM folder
WHERE id = $1
`

func (q *Queries) DeleteFolder(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFolder, id)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/folder.sql.go/DeleteFolder() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return err
	}
	return nil
}
