-- +goose Up
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT ENCODE(SHA256(RANDOM()::TEXT::BYTEA), 'hex');

-- +goose Down 
ALTER TABLE users DROP COLUMN api_key;

