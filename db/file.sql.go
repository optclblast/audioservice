package db

import (
	"context"
	"fmt"
	"time"

	"github.com/optclblast/filetagger/logger"
)

// operations with files
type CreateFileParams struct {
	Id     int64  `json:"id"`
	Owner  int64  `json:"owner"`
	Parent int64  `json:"parent"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Tag    string `json:"tag"`
}

const createFile = `-- name: CreateFile :one
INSERT INTO file (
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

func (q *Queries) CreateFile(ctx context.Context, arg CreateFileParams) (File, error) {
	row := q.db.QueryRowContext(ctx, createFile, arg.Id, arg.Owner, arg.Parent, arg.Name, time.Now(), arg.Path, arg.Tag)
	var i File
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
			Location: "db/sqlc/file.sql.go/CreateFile() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const getFile = `-- name: GetFile :one
SELECT id, owner, parent, name, path, tag FROM file
WHERE id = $1 AND owner = $2 LIMIT 1
`

func (q *Queries) GetFile(ctx context.Context, id int64, owner int64) (File, error) {
	row := q.db.QueryRowContext(ctx, getFile, id, owner)
	var i File
	err := row.Scan(
		&i.Id,
		&i.Owner,
		&i.Parent,
		&i.Name,
		&i.Path,
		&i.Tag,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/file.sql.go/GetFile() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const listFiles = `--name: ListFiles :many
SELECT id, owner, parent, name, path, tag FROM file
WHERE owner = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListFilesParams struct {
	Limit  int32 `json:"id"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFiles(ctx context.Context, arg ListFilesParams, owner int64) ([]File, error) {
	rows, err := q.db.QueryContext(ctx, listFiles, arg.Limit, arg.Offset, owner)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/file.sql.go/ListFiles() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.Id,
			&i.Owner,
			&i.Parent,
			&i.Name,
			&i.Path,
			&i.Tag,
		); err != nil {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "db/sqlc/file.sql.go/ListFiles() SCOPE",
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
			Location: "db/sqlc/file.sql.go/ListFiles() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}
	if err := rows.Err(); err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/file.sql.go/ListFiles() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	return items, nil
}

const updateFile = `--name: UpdateFile :one
UPDATE file
SET name = $2, path = $3, tag = $4
WHERE id = $1
RETURNING id, name, path, tag
`

type UpdateFileParams struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Tag  string `json:"tag"`
}

func (q *Queries) UpdateFile(ctx context.Context, arg UpdateFileParams) (Folder, error) {
	row := q.db.QueryRowContext(ctx, updateFile, arg.Id, arg.Name, arg.Path, arg.Tag)
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
			Location: "db/sqlc/file.sql.go/UpdateFiles() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const deleteFile = `--name: DeleteFile :exec
DELETE FROM file
WHERE id = $1
`

func (q *Queries) DeleteFile(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFile, id)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/file.sql.go/DeleteFile() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return err
	}
	return nil
}
