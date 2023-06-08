package db

import (
	"context"
	"fmt"
	"time"

	"github.com/optclblast/filetagger/logger"
)

// operations with files
type CreateRSAKeyParams struct {
	Owner int64  `json:"owner"`
	Key   string `json:"key"`
}

const createRSAKey = `-- name: CreateRSAKey :one
INSERT INTO rsakeys (
	owner,
	key
) VALUES (
	$1, $2
) RETURNING id, owner, key
`

func (q *Queries) CreateRSAKey(ctx context.Context, arg CreateRSAKeyParams) (RSAKey, error) {
	row := q.db.QueryRowContext(ctx, createRSAKey, arg.Owner, arg.Key)
	var i RSAKey
	err := row.Scan(
		&i.Id,
		&i.Owner,
		&i.Key,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/rsakeys.sql.go/CreateRSAKey() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const getRSAKey = `-- name: GetRSAKey :one
SELECT key FROM rsakeys
WHERE owner = $1 
ORDER BY id
LIMIT 1
`

func (q *Queries) GetRSAKey(ctx context.Context, owner int64) (RSAKey, error) {
	row := q.db.QueryRowContext(ctx, getRSAKey, owner)
	var i RSAKey
	err := row.Scan(
		&i.Key,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/rsakeys.sql.go/GetRSAKey() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const updateRSAKey = `--name: UpdateRSAKey :one
UPDATE rsakeys
SET key = $2
WHERE owner = $1
RETURNING owner, key
`

type UpdateRSAKeyParams struct {
	Owner int64  `json:"owner"`
	Key   string `json:"key"`
}

func (q *Queries) UpdateRSAKey(ctx context.Context, arg UpdateRSAKeyParams) (RSAKey, error) {
	row := q.db.QueryRowContext(ctx, updateRSAKey, arg.Owner, arg.Key)
	var i RSAKey
	err := row.Scan(
		&i.Owner,
		&i.Key,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/rsakeys.sql.go/UpdateRSAKey() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const deleteRSAKey = `--name: DeleteRSAKey :exec
DELETE FROM rsakeys
WHERE owner = $1
`

func (q *Queries) DeleteRSAKey(ctx context.Context, owner int64) error {
	_, err := q.db.ExecContext(ctx, deleteRSAKey, owner)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/rsakeys.sql.go/DeleteRSAKey() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return err
	}
	return nil
}
