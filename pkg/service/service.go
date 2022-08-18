package service

import (
	"boosters/api/pkg/repository"
	"boosters/api/post"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service struct {
	Posts
}
type Posts interface {
	CreatePost(post post.Post) (int, error)
	GetAllPosts() ([]post.Post, error)
	GetById(id string) (post.Post, error)
	UpdatePost(post post.Post, id string) error
	DeletePost(id string) error
}

func NewService(repo repository.Repository) *Service {
	return &Service{Posts: NewPostsService(repo.Posts)}
}
