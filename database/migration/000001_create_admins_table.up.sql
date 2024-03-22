CREATE TABLE admins (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
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

CREATE UNIQUE INDEX admins_email_unique ON admins(email);

CREATE INDEX idx_admins_email ON admins(email);

CREATE INDEX idx_admins_created_by ON admins(created_by);

CREATE INDEX idx_admins_updated_by ON admins(updated_by);