-- name: CreatePost :one
INSERT INTO posts (
	id, created_at, updated_at, title, url, description, published_at, feed_id
) VALUES ($1, NOW(), NOW(), $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPostsByUser :many
SELECT *
FROM posts p 
LEFT JOIN feeds f
ON p.feed_id = f.id
WHERE f.user_id = $1
ORDER BY p.created_at DESC
LIMIT $2;

