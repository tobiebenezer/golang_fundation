package store

import (
	"context"
	"database/sql"
	"errors"

	"social/update/internal/domain"
)

type CommentStore struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) domain.CommentRepository {
	return &CommentStore{db: db}
}

func (c *CommentStore) Insert(ctx context.Context, comment *domain.Comment) error {
	query := `
	INSERT INTO comments(post_id, user_id, content) VALUES($1, $2, $3) RETURNING id`
	row := c.db.QueryRowContext(ctx, query, comment.PostID, comment.UserID, comment.Content)

	if err := row.Scan(&comment.ID); err != nil {
		return err
	}

	return nil
}

func (c *CommentStore) GetByID(ctx context.Context, id int) (*domain.Comment, error) {
	query := `
	SELECT * FROM comments 
	WHERE id = $1 AND deleted_at IS NULL`
	row := c.db.QueryRowContext(ctx, query, id)

	var comment domain.Comment

	if err := row.Scan(
		&comment.ID,
		&comment.PostID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&comment.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}


	return &comment, nil
}

func (c *CommentStore) Update(ctx context.Context, id int, comment *domain.Comment) (*domain.Comment, error) {
	query := `
	UPDATE comments 
	SET content = $2,
		post_id = COALESCE($3, post_id),
		user_id = COALESCE($4, user_id),
		updated_at = CURRENT_TIMESTAMP
	WHERE id = $1 AND deleted_at IS NULL
	RETURNING id, post_id, user_id, content, created_at, updated_at, deleted_at`
	
	row := c.db.QueryRowContext(ctx, query, id, comment.Content, comment.PostID, comment.UserID)
	
	var updatedComment domain.Comment
	if err := row.Scan(
		&updatedComment.ID,
		&updatedComment.PostID,
		&updatedComment.UserID,
		&updatedComment.Content,
		&updatedComment.CreatedAt,
		&updatedComment.UpdatedAt,
		&updatedComment.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &updatedComment, nil
}

func (c *CommentStore) Delete(ctx context.Context, id int) error {
	query := `
	UPDATE comments 
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE id = $1 AND deleted_at IS NULL`
	
	result, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return domain.ErrNotFound
	}
	
	return nil
}

func (c *CommentStore) ListByPostID(ctx context.Context, postID int, limit, offset int) ([]domain.Comment, error) {
	query := `
	SELECT id, post_id, user_id, content, created_at, updated_at, deleted_at
	FROM comments
	WHERE post_id = $1 AND deleted_at IS NULL
	ORDER BY id DESC
	LIMIT $2 OFFSET $3
	`
	
	rows, err := c.db.QueryContext(ctx, query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var comments []domain.Comment
	for rows.Next() {
		var comment domain.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.DeletedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return comments, nil
}

func (c *CommentStore) ListByUserID(ctx context.Context, userID int, limit, offset int) ([]domain.Comment, error) {
	query := `
	SELECT id, post_id, user_id, content, created_at, updated_at, deleted_at
	FROM comments
	WHERE user_id = $1 AND deleted_at IS NULL
	ORDER BY id DESC
	LIMIT $2 OFFSET $3
	`
	
	rows, err := c.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var comments []domain.Comment
	for rows.Next() {
		var comment domain.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.DeletedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return comments, nil
}

func (c *CommentStore) CountByPostID(ctx context.Context, postID int) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM comments
	WHERE post_id = $1 AND deleted_at IS NULL
	`
	
	var count int
	err := c.db.QueryRowContext(ctx, query, postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	
	return count, nil
}

func (c *CommentStore) CountByUserID(ctx context.Context, userID int) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM comments
	WHERE user_id = $1 AND deleted_at IS NULL
	`
	
	var count int
	err := c.db.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	
	return count, nil
}
