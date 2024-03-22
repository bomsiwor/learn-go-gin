CREATE TABLE carts (
    id BIGSERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE
    SET
        NULL,
        product_id INT REFERENCES products(id) ON DELETE
    SET
        NULL,
        quantity INT NOT NULL DEFAULT 1,
        is_checked BOOLEAN NOT NULL DEFAULT TRUE,
        created_by INT REFERENCES users(id) ON DELETE
    SET
        NULL,
        updated_by INT REFERENCES users(id) ON DELETE
    SET
        NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP,
        deleted_at TIMESTAMP
);

CREATE INDEX idx_carts_user_id ON carts(user_id);

CREATE INDEX idx_carts_product_id ON carts(product_id);

CREATE INDEX idx_carts_created_by ON carts(created_by);

CREATE INDEX idx_carts_updated_by ON carts(updated_by);