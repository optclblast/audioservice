package db

import (
	"context"
	"fmt"
	"time"

	"github.com/optclblast/filetagger/logger"
)

type CreateAccountParams struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//user data operations

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (
	login,
	password
) VALUES (
	$1, $2
) RETURNING id, login, password, created_at
`

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Login, arg.Password)
	var i Account
	err := row.Scan(
		&i.Id,
		&i.Login,
		&i.Password,
		&i.CreatedAt,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/account.sql.go/CreateAccount() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const getAccount = `--name: GetAccount :one
SELECT id, login, password, created_at FROM account
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.Id,
		&i.Login,
		&i.Password,
		&i.CreatedAt,
	)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "db/sqlc/account.sql.go/GetAccount() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return i, err
}

const listAccounts = `--name: ListAccounts :many
SELECT id, login, password, created_at FROM account
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/account.sql.go/ListAccounts() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}
	defer rows.Close()

	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.Id,
			&i.Login,
			&i.Password,
			&i.CreatedAt,
		); err != nil {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "db/sqlc/account.sql.go/ListAccounts() SCOPE",
				Content:  fmt.Sprint(err),
			})
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/account.sql.go/ListAccounts() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}
	if err := rows.Err(); err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/account.sql.go/ListAccounts() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}
	return items, nil
}

const updateAccount = `--name: UpdateAccount :one
UPDATE account
SET login = $2
WHERE id = &1
RETURNING id, login, created_at
`

type UpdateAccountParams struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.Id, arg.Login)

	var i Account
	err := row.Scan(
		&i.Id,
		&i.Login,
		&i.Password,
		&i.CreatedAt,
	)

	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/account.sql.go/UpdateAccount() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}

	return i, err
}

const deleteAccount = `--name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "db/sqlc/account.sql.go/DeleteAccount() SCOPE",
			Content:  fmt.Sprint(err),
		})
	}
	return err
}
