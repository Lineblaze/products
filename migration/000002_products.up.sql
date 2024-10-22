CREATE TABLE products
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    "text" TEXT,
    category_id BIGINT REFERENCES product_categories(id)
);
