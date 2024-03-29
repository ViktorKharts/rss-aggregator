-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsByLastFetchedAt :many
SELECT * FROM feeds
ORDER BY last_fetched_at DESC
LIMIT $1;

-- name: UpdateFeedLastFetchedAt :many 
UPDATE feeds
SET updated_at = NOW(), last_fetched_at = NOW() 
WHERE id = $1
RETURNING *;

