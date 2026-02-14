-- +goose Up
CREATE TABLE product (
	id SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
	description VARCHAR NOT NULL,
	price INTEGER NOT NULL,
	quantity INTEGER NOT NULL DEFAULT 0,
	image_url VARCHAR
);

CREATE INDEX idx_product_id ON product (id);

-- +goose Down
DROP INDEX IF EXISTS idx_product_id;
DROP TABLE IF EXISTS product;
