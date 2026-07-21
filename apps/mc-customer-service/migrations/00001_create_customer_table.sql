-- +goose Up
CREATE TABLE customer (
	id VARCHAR PRIMARY KEY,
	name VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	password VARCHAR NOT NULL
);

CREATE INDEX idx_customer_id ON customer (id);
CREATE INDEX idx_customer_email ON customer (email);

-- +goose Down
DROP INDEX IF EXISTS idx_customer_email;
DROP INDEX IF EXISTS idx_customer_id;
DROP TABLE IF EXISTS customer;
