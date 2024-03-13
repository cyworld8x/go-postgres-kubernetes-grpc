-- name: GetFollowers :many
SELECT * FROM follows
WHERE followed_user_id = $1 LIMIT 1;

-- name: FollowUser :one
INSERT INTO follows (
  following_user_id, followed_user_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UnfollowUser :exec
DELETE FROM follows
WHERE followed_user_id = $1 AND following_user_id = $2;