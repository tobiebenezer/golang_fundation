package store

import (
	"context"
	"database/sql"
	"errors"

	"social/update/internal/domain"
)

type UserStore struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserStore{db: db}
}

func (p *UserStore) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (username, email, password) 
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at
	`

	row := p.db.QueryRowContext(ctx, query,
		user.Username,
		user.Email,
		[]byte(user.Password), // DB expects bytea
	)

	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserStore) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at 
	FROM users 
	WHERE id = $1 AND deleted_at IS NULL`

	var user domain.User
	var passwordHash []byte

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&passwordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	user.Password = string(passwordHash)
	return &user, nil
}

func (p *UserStore) Update(ctx context.Context, user *domain.User) error {
	query := `UPDATE users 
	SET username = $1, email = $2, updated_at = CURRENT_TIMESTAMP 
	WHERE id = $3 AND deleted_at IS NULL`

	_, err := p.db.ExecContext(ctx, query, user.Username, user.Email, user.ID)
	return err
}

func (p *UserStore) Delete(ctx context.Context, id int) error {
	query := `UPDATE users 
	SET deleted_at = CURRENT_TIMESTAMP 
	WHERE id = $1`

	_, err := p.db.ExecContext(ctx, query, id)
	return err
}

func (p *UserStore) List(ctx context.Context, limit, offset int) ([]domain.User, error) {
	query := `SELECT id, username, email, created_at, updated_at 
	FROM users 
	WHERE deleted_at IS NULL 
	ORDER BY id ASC 
	LIMIT $1 OFFSET $2`

	rows, err := p.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []domain.User{}
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}