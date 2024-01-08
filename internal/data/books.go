package data

import (
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Isbn      int32     `json:"isbn"`
	Year      int32     `json:"year,omitempty"`
	Genres    []string  `json:"genres,omitempty"`

	Version int32 // The version number starts at 1 and will be incremented each
	// time the book information is updated
}
