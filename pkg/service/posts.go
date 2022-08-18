package service

import (
	"boosters/api/pkg/repository"
	"boosters/api/post"
)

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(repo repository.Posts) *PostsService {
	return &PostsService{repo: repo}
}

func (s *PostsService) CreatePost(post post.Post) (int, error) {
	return s.repo.CreatePost(post)
}

func (s *PostsService) GetAllPosts() ([]post.Post, error) {
	return s.repo.GetAllPosts()
}
func (s *PostsService) GetById(id string) (post.Post, error) {
	return s.repo.GetById(id)
}
func (s *PostsService) UpdatePost(post post.Post, id string) error {
	return s.repo.UpdatePost(post, id)
}
func (s *PostsService) DeletePost(id string) error {
	return s.repo.DeletePost(id)
}
