package domain

import "context"

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, limit, offset int) ([]User, error)
}

type UserService interface {
	Register(ctx context.Context, username, email, password string) (*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, id int, username, email *string) (*User, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, limit, offset int) ([]User, error)
}
