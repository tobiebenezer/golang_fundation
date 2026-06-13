package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"social/update/internal/domain"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, username, email, password string) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) Update(ctx context.Context, id int, username, email *string) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if username != nil {
		user.Username = *username
	}
	if email != nil {
		user.Email = *email
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Delete(ctx context.Context, id int) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s *userService) List(ctx context.Context, limit, offset int) ([]domain.User, error) {
	return s.repo.List(ctx, limit, offset)
}
