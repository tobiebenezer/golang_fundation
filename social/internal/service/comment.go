package service

import (
	"context"
	"time"

	"social/update/internal/domain"
)

type commentService struct {
	commentRepo domain.CommentRepository
	postRepo    domain.PostRepository
	userRepo    domain.UserRepository
}

func NewCommentService(
	commentRepo domain.CommentRepository,
	postRepo domain.PostRepository,
	userRepo domain.UserRepository,
) domain.CommentService {
	return &commentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
		userRepo:    userRepo,
	}
}

func (s *commentService) CreateComment(ctx context.Context, comment *domain.Comment) error {
	// Validate post exists (GetByID will fail if deleted/not found)
	_, err := s.postRepo.GetByID(ctx, comment.PostID)
	if err != nil {
		return err
	}

	// Validate user exists (GetByID will fail if deleted/not found)
	_, err = s.userRepo.GetByID(ctx, comment.UserID)
	if err != nil {
		return err
	}

	// Create comment
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	return s.commentRepo.Insert(ctx, comment)
}

func (s *commentService) GetCommentByID(ctx context.Context, id int) (*domain.Comment, error) {
	return s.commentRepo.GetByID(ctx, id)
}

func (s *commentService) UpdateComment(ctx context.Context, id int, comment *domain.Comment) (*domain.Comment, error) {
	// Validate comment exists
	existing, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if comment.Content != "" {
		existing.Content = comment.Content
	}
	existing.UpdatedAt = time.Now()

	return s.commentRepo.Update(ctx, id, existing)
}

func (s *commentService) DeleteComment(ctx context.Context, id int) error {
	// Validate comment exists
	_, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.commentRepo.Delete(ctx, id)
}

func (s *commentService) ListCommentsByPostID(ctx context.Context, postID int, limit, offset int) ([]domain.Comment, error) {
	_, err := s.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return s.commentRepo.ListByPostID(ctx, postID, limit, offset)
}

func (s *commentService) ListCommentsByUserID(ctx context.Context, userID int, limit, offset int) ([]domain.Comment, error) {
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return s.commentRepo.ListByUserID(ctx, userID, limit, offset)
}

func (s *commentService) CountCommentsByPostID(ctx context.Context, postID int) (int, error) {
	_, err := s.postRepo.GetByID(ctx, postID)
	if err != nil {
		return 0, err
	}

	return s.commentRepo.CountByPostID(ctx, postID)
}

func (s *commentService) CountCommentsByUserID(ctx context.Context, userID int) (int, error) {
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return 0, err
	}

	return s.commentRepo.CountByUserID(ctx, userID)
}
