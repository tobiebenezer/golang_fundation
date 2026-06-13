package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"social/update/internal/domain"
)

type PostStore struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) domain.PostRepository {
	return &PostStore{db: db}
}

func (p *PostStore) Create(ctx context.Context, post *domain.Post) error {
	query := `INSERT INTO posts (title, content, user_id, tags) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at
	`

	row := p.db.QueryRowContext(ctx, query,	
		post.Title,
		post.Content,
		post.UserID,
		pq.Array(post.Tags),
	)

	err := row.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostStore) GetByID(ctx context.Context, id int) (*domain.Post, error) {
	query := `SELECT id, title, content, user_id, tags, created_at, updated_at 
	FROM posts 
	WHERE id = $1 AND deleted_at IS NULL`

	var post domain.Post
	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &post, nil
}

func (p *PostStore) Update(ctx context.Context, post *domain.Post) error {
	query := `UPDATE posts 
	SET title = $1, content = $2, tags = $3, updated_at = CURRENT_TIMESTAMP 
	WHERE id = $4 AND deleted_at IS NULL`

	_, err := p.db.ExecContext(ctx, query, post.Title, post.Content, pq.Array(post.Tags), post.ID)
	return err
}

func (p *PostStore) Delete(ctx context.Context, id int) error {
	query := `UPDATE posts 
	SET deleted_at = CURRENT_TIMESTAMP 
	WHERE id = $1`

	_, err := p.db.ExecContext(ctx, query, id)
	return err
}

func (p *PostStore) List(ctx context.Context, limit, offset int) ([]domain.Post, error) {
	query := `SELECT id, title, content, user_id, tags, created_at, updated_at 
	FROM posts 
	WHERE deleted_at IS NULL 
	ORDER BY id DESC 
	LIMIT $1 OFFSET $2`

	rows, err := p.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []domain.Post{}
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.UserID,
			pq.Array(&post.Tags),
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}