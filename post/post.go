package post

import "time"

type Post struct {
	Id         int       `json:"id" db:"id"`
	Title      string    `json:"title" db:"title" validate:"required"`
	Content    string    `json:"content" db:"content" validate:"required"`
	Created_at time.Time `json:"create_at" db:"create_at"`
	Updated_at time.Time `json:"update_at" db:"update_at"`
}
