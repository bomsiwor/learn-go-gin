CREATE TABLE oauth_access_tokens (
    id BIGSERIAL PRIMARY KEY,
    oauth_client_id INT REFERENCES oauth_clients(id) ON DELETE
    SET
        NULL,
        user_id INT NOT NULL,
        token VARCHAR(255),
        scope VARCHAR(255),
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

CREATE UNIQUE INDEX oauth_access_tokens_token_unique ON oauth_access_tokens(token);

CREATE INDEX idx_oauth_access_tokens_token ON oauth_access_tokens(token);

CREATE INDEX idx_oauth_access_oauth_client_id ON oauth_access_tokens(oauth_client_id);

CREATE INDEX idx_oauth_access_tokens_created_by ON oauth_access_tokens(created_by);

CREATE INDEX idx_oauth_access_tokens_updated_by ON oauth_access_tokens(updated_by);