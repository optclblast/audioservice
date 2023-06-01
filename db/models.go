package db

import (
	"database/sql"
	"time"
)

type Account struct {
	Id        int64     `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

type RSAKey struct {
	Id    int64  `json:"id"`
	Owner int64  `json:"owner"`
	Key   string `json:"key"`
}

type Folder struct {
	Id        int64         `json:"id"`
	Owner     int64         `json:"owner"`
	Parent    sql.NullInt64 `json:"parent"`
	Name      string        `json:"name"`
	CreatedAt time.Time     `json:"createdat"`
	Path      string        `json:"content"`
	Tag       string        `json:"tag"`
}

type File struct {
	Id        int64         `json:"id"`
	Owner     int64         `json:"owner"`
	Parent    sql.NullInt64 `json:"parent"`
	Name      string        `json:"name"`
	CreatedAt time.Time     `json:"createdat"`
	Path      string        `json:"content"`
	Tag       string        `json:"tag"`
}
