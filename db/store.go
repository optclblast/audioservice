package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type RegisterAccountTxParams struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

type RegisterAccountTxResult struct {
	Account Account `json:"account"`
	Key     RSAKey  `json:"rsakey"`
}

func (store *Store) TransferTx(ctx context.Context, arg RegisterAccountTxParams) (RegisterAccountTxResult, error) {
	var result RegisterAccountTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Account, err = q.CreateAccount(ctx, CreateAccountParams{
			arg.Login,
			arg.Password,
		})
		if err != nil {
			return err
		}

		result.Key, err = q.CreateRSAKey(ctx, CreateRSAKeyParams{
			result.Account.Id,
			arg.Key,
		})
		if err != nil {
			return err
		}
		//TODO

		return nil
	})

	return result, err
}
