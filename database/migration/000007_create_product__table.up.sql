CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    product_category_id INT REFERENCES product_categories(id) ON DELETE
    SET
        NULL,
        title VARCHAR(255) NOT NULL,
        image VARCHAR(255),
        video VARCHAR(255),
        description VARCHAR(255),
        is_highlighted BOOLEAN DEFAULT FALSE NOT NULL,
        price INT NOT NULL,
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

CREATE INDEX idx_products_created_by ON products(created_by);

CREATE INDEX idx_products_updated_by ON products(updated_by);

CREATE INDEX idx_products_product_category_id ON products(product_category_id);