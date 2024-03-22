CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE
    SET
        NULL,
        discount_id INT REFERENCES discounts(id) ON DELETE
    SET
        NULL,
        checkout_link VARCHAR(255),
        external_id VARCHAR(255),
        price INT NOT NULL,
        total_price INT NOT NULL,
        status VARCHAR(255) NOT NULL,
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

CREATE INDEX idx_orders_user_id ON orders(user_id);

CREATE INDEX idx_orders_discount_id ON orders(discount_id);

CREATE INDEX idx_orders_created_by ON orders(created_by);

CREATE INDEX idx_orders_updated_by ON orders(updated_by);