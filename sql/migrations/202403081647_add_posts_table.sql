-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title VARCHAR NOT NULL,
	url VARCHAR UNIQUE NOT NULL,
	description VARCHAR,
	published_at TIMESTAMP,
	feed_id UUID NOT NULL,
	FOREIGN KEY(feed_id) 
	REFERENCES feeds(id)
	ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS posts;
