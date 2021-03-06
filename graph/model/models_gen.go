// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Album struct {
	Name  string  `json:"name"`
	Songs []*Song `json:"songs"`
}

type Artist struct {
	Name   string   `json:"name"`
	Albums []*Album `json:"albums"`
}

type CreatePlaylist struct {
	Name string `json:"name"`
}

type Playlist struct {
	Name  string  `json:"name"`
	Songs []*Song `json:"songs"`
}

type Song struct {
	ID           string  `json:"id"`
	Artist       string  `json:"artist"`
	Duration     string  `json:"duration"`
	URI          *string `json:"uri"`
	LastModified *string `json:"lastModified"`
	Title        string  `json:"title"`
	Album        string  `json:"album"`
	Track        *string `json:"track"`
}
