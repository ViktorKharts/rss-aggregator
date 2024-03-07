-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsByLastFetchedAt :many
SELECT * FROM feeds
ORDER BY last_fetched_at DESC;

-- name: UpdateFeedLastFetchedAt :exec
UPDATE feeds
SET updated_at = $2, last_fetched_at = $2 
WHERE id = $1;

