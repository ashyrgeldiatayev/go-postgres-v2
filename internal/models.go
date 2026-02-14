package internal

import "time"

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Books struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
