package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      string    `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty,string"` // int could be string
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
