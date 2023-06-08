package db

import (
	"database/sql"
	"time"
)

type Account struct {
	Id       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Artist struct {
	Id   int64  `json:"id"`
	Name int64  `json:"name"`
	Bio  string `json:"bio"`
}

type Collection struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Author      int64         `json:"author"`
	FtAuthors   sql.NullInt64 `json:"ft_authors"`
	Type        string        `json:"type"`
	Discription string        `json:"discription"`
	Lenght      string        `json:"lenght"`
	Label       string        `json:"label"`
	Date        time.Time     `json:"date"`
}

type Track struct {
	Id        int64         `json:"id"`
	Name      string        `json:"name"`
	Author    int64         `json:"author"`
	FtAuthors sql.NullInt64 `json:"ft_authors"`
	Type      string
}
