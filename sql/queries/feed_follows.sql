-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    inserted.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted
INNER JOIN users
    ON  users.id = inserted.user_id
INNER JOIN feeds
    ON  feeds.id = inserted.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows
INNER JOIN users
    ON  users.id = feed_follows.user_id
INNER JOIN feeds
    ON  feeds.id = feed_follows.feed_id
WHERE
    users.id = $1
ORDER BY
    feeds.name;