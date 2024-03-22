CREATE TABLE class_rooms (
    id BIGSERIAL PRIMARY KEY,
    user_id INT,
    product_id INT,
    created_by INT,
    updated_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT FK_class_rooms_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE
    SET
        NULL,
        CONSTRAINT FK_class_rooms_product_id FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE
    SET
        NULL,
        CONSTRAINT FK_class_rooms_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE
    SET
        NULL,
        CONSTRAINT FK_class_rooms_updated_by FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE
    SET
        NULL
);

CREATE INDEX idx_class_rooms_user_id ON class_rooms(user_id);

CREATE INDEX idx_class_rooms_product_id ON class_rooms(product_id);

CREATE INDEX idx_class_rooms_created_by ON class_rooms(created_by);

CREATE INDEX idx_class_rooms_updated_by ON class_rooms(updated_by);