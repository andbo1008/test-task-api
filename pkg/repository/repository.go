package repository

import (
	"boosters/api/post"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Posts
}

type Posts interface {
	CreatePost(post post.Post) (int, error)
	GetAllPosts() ([]post.Post, error)
	GetById(id string) (post.Post, error)
	UpdatePost(post post.Post, id string) error
	DeletePost(id string) error
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Posts: NewPosts(db),
	}
}
