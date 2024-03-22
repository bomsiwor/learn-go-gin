CREATE TABLE product_categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
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

CREATE INDEX idx_product_categories_created_by ON product_categories(created_by);

CREATE INDEX idx_product_categories_updated_by ON product_categories(updated_by);