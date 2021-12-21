package data

import "time"

type Movie struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      string
	Runtime   int32
	Genres    []string
	Version   int32
}
