CREATE TABLE discounts (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) UNIQUE NOT NULL,
    quantity INT NOT NULL,
    remaining_quantity INT NOT NULL,
    type VARCHAR(255) NOT NULL,
    value INT NOT NULL,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
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

CREATE INDEX idx_discounts_created_by ON discounts(created_by);

CREATE INDEX idx_discounts_updated_by ON discounts(updated_by);