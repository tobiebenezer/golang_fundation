package service

import (
	"context"

	"social/update/internal/domain"
)

type postService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) domain.PostService {
	return &postService{repo: repo}
}

func (s *postService) Create(ctx context.Context, post *domain.Post) error {
	return s.repo.Create(ctx, post)
}

func (s *postService) GetByID(ctx context.Context, id int) (*domain.Post, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *postService) Update(ctx context.Context, id int, title, content *string, tags []string) (*domain.Post, error) {
	post, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if title != nil {
		post.Title = *title
	}
	if content != nil {
		post.Content = *content
	}
	if tags != nil {
		post.Tags = tags
	}

	if err := s.repo.Update(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) Delete(ctx context.Context, id int) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s *postService) List(ctx context.Context, limit, offset int) ([]domain.Post, error) {
	return s.repo.List(ctx, limit, offset)
}
