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
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type Collection struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Author      int64         `json:"artist"`
	FtAuthors   sql.NullInt64 `json:"ft_artists"`
	Type        string        `json:"type"`
	Discription string        `json:"discription"`
	Lenght      string        `json:"lenght"`
	Label       string        `json:"label"`
	Date        time.Time     `json:"date"`
}

type Track struct {
	Id        int64         `json:"id"`
	Name      string        `json:"name"`
	Artist    int64         `json:"artist"`
	FtArtists sql.NullInt64 `json:"ft_artists"`
	Album     int64         `json:"album"`
	Location  string        `json:"location"`
}

type Playlist struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Owner       int64     `json:"owner"`
	Discription string    `json:"discription"`
	Lenght      string    `json:"lenght"`
	Date        time.Time `json:"date"`
	UploadDate  time.Time `json:"upload_date"`
}

type UserLikedTracks struct {
	User   int64 `json:"user"`
	Tracks int64 `json:"tracks"`
}

type UserLikedCollections struct {
	User        int64 `json:"user"`
	Collections int64 `json:"collections"`
}

type UserLikesPlaylists struct {
	User     int64 `json:"user"`
	Playlist int64 `json:"playlist"`
}
