CREATE TABLE forgot_passwords (
    id BIGSERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE
    SET
        NULL,
        valid BOOLEAN NOT NULL DEFAULT TRUE,
        code VARCHAR(255) NOT NULL,
        expired_at TIMESTAMP,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP,
        deleted_at TIMESTAMP
);

CREATE INDEX idx_forgot_passwords_user_id ON forgot_passwords(user_id);