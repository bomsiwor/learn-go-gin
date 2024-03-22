CREATE TABLE order_details (
    id BIGSERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE
    SET
        NULL,
        product_id INT REFERENCES products(id) ON DELETE
    SET
        NULL,
        price INT NOT NULL,
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

CREATE INDEX idx_order_details_order_id ON order_details(order_id);

CREATE INDEX idx_order_details_product_id ON order_details(product_id);

CREATE INDEX idx_order_details_created_by ON order_details(created_by);

CREATE INDEX idx_order_details_updated_by ON order_details(updated_by);