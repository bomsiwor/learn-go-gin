CREATE TABLE oauth_refresh_tokens (
    id BIGSERIAL PRIMARY KEY,
    oauth_access_token_id INT REFERENCES oauth_access_tokens(id) ON DELETE
    SET
        NULL,
        user_id INT NOT NULL,
        token VARCHAR(255) UNIQUE,
        expired_at TIMESTAMP,
        created_by INT REFERENCES admins(id) ON DELETE
    SET
        NULL,
        updated_by INT REFERENCES admins(id) ON DELETE
    SET
        NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP,
        deleted_at TIMESTAMP
);

CREATE INDEX idx_oauth_refresh_tokens_oauth_access_token_id ON oauth_refresh_tokens(oauth_access_token_id);

CREATE INDEX idx_oauth_refresh_tokens_token ON oauth_refresh_tokens(token);

CREATE INDEX idx_oauth_refresh_tokens_created_by ON oauth_refresh_tokens(created_by);

CREATE INDEX idx_oauth_refresh_tokens_updated_by ON oauth_refresh_tokens(updated_by);