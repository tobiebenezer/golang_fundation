package domain

import (
	"context"
	"time"
)


type Comment struct{
	ID int `json:"id"`
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CommentRepository interface {
	Insert(ctx context.Context, comment *Comment) error
	GetByID(ctx context.Context, id int) (*Comment, error)
	Update(ctx context.Context, id int, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id int) error
	ListByPostID(ctx context.Context, postID int, limit, offset int) ([]Comment, error)
	ListByUserID(ctx context.Context, userID int, limit, offset int) ([]Comment, error)
	CountByPostID(ctx context.Context, postID int) (int, error)
	CountByUserID(ctx context.Context, userID int) (int, error)
}

func (c *Comment) TableName() string {
	return "comments"
}

type CommentService interface {
	CreateComment(ctx context.Context, comment *Comment) error
	GetCommentByID(ctx context.Context, id int) (*Comment, error)
	UpdateComment(ctx context.Context, id int, comment *Comment) (*Comment, error)
	DeleteComment(ctx context.Context, id int) error
	ListCommentsByPostID(ctx context.Context, postID int, limit, offset int) ([]Comment, error)
	ListCommentsByUserID(ctx context.Context, userID int, limit, offset int) ([]Comment, error)
	CountCommentsByPostID(ctx context.Context, postID int) (int, error)
	CountCommentsByUserID(ctx context.Context, userID int) (int, error)
}
