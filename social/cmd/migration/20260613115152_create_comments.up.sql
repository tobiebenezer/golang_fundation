CREATE TABLE IF NOT EXISTS comments(
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    CONSTRAINT fk_posts_comments
    FOREIGN KEY (post_id)
    REFERENCES posts(id)
    ON DELETE CASCADE,

    CONSTRAINT fk_users_comments
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

-- Indexes
CREATE INDEX idx_comments_post_id ON comments(post_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_comments_user_id ON comments(user_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_comments_created_at ON comments(created_at DESC) WHERE deleted_at IS NULL;