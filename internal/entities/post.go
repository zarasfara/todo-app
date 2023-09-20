package entities

import "time"

type Post struct {
	ID      int       `json:"id"`
	Title   string    `json:title`
	Content string    `json:content`
	Created time.Time `json:created_at`
}
