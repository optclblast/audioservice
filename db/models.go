package db

import "time"

type Account struct {
	Id        int64     `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

type RSAKey struct {
	Owner int64  `json:"owner"`
	Key   string `json:"key"`
}

type Folder struct {
	Id          int64     `json:"id"`
	Owner       int64     `json:"owner"`
	Parent      int64     `json:"parent"`
	Name        string    `json:"name"`
	AccessLevel string    `json:"accesslevel"`
	CreatedAt   time.Time `json:"createdat"`
	Content     int64     `json:"content"`
	Tag         string    `json:"tag"`
}

type File struct {
	Id      int64  `json:"id"`
	Owner   int64  `json:"owner"`
	Parent  int64  `json:"parent"`
	Name    string `json:"name"`
	Content int64  `json:"content"`
	Tag     string `json:"tag"`
}
