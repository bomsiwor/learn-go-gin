CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    code_verified VARCHAR(255) NOT NULL,
    email_verified_at TIMESTAMP,
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

CREATE UNIQUE INDEX users_email_unique ON users(email);

CREATE INDEX idx_users_email ON users(email);

CREATE INDEX idx_users_created_by ON users(created_by);

CREATE INDEX idx_users_updated_by ON users(updated_by);