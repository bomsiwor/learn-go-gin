CREATE TABLE oauth_clients (
    id BIGSERIAL PRIMARY KEY,
    client_id VARCHAR(255) NOT NULL,
    client_secret VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    redirect VARCHAR(255),
    description VARCHAR(255),
    scope VARCHAR(255),
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

CREATE UNIQUE INDEX oauth_clients_client_id_unique ON oauth_clients(client_id);

CREATE INDEX idx_oauth_clients_client_id ON oauth_clients(client_id);

CREATE INDEX idx_oauth_client_created_by ON oauth_clients(created_by);

CREATE INDEX idx_oauth_client_updated_by ON oauth_clients(updated_by);