package domain

import (
	"context"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	GetByID(ctx context.Context, id int) (*Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, limit, offset int) ([]Post, error)
}

type PostService interface {
	Create(ctx context.Context, post *Post) error
	GetByID(ctx context.Context, id int) (*Post, error)
	Update(ctx context.Context, id int, title, content *string, tags []string) (*Post, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, limit, offset int) ([]Post, error)
}
