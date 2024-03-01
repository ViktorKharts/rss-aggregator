-- +goose Up
CREATE TABLE IF NOT EXISTS users (
	id VARCHAR PRIMARY KEY,
	name VARCHAR NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
