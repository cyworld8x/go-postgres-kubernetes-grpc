// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: follow.sql

package db

import (
	"context"
)

const followUser = `-- name: FollowUser :one
INSERT INTO follows (
  following_user_id, followed_user_id
) VALUES (
  $1, $2
)
RETURNING following_user_id, followed_user_id, created_at
`

type FollowUserParams struct {
	FollowingUserID int32 `json:"following_user_id"`
	FollowedUserID  int32 `json:"followed_user_id"`
}

func (q *Queries) FollowUser(ctx context.Context, arg FollowUserParams) (Follow, error) {
	row := q.db.QueryRow(ctx, followUser, arg.FollowingUserID, arg.FollowedUserID)
	var i Follow
	err := row.Scan(&i.FollowingUserID, &i.FollowedUserID, &i.CreatedAt)
	return i, err
}

const getFollowers = `-- name: GetFollowers :many
SELECT following_user_id, followed_user_id, created_at FROM follows
WHERE followed_user_id = $1 LIMIT 1
`

func (q *Queries) GetFollowers(ctx context.Context, followedUserID int32) ([]Follow, error) {
	rows, err := q.db.Query(ctx, getFollowers, followedUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Follow{}
	for rows.Next() {
		var i Follow
		if err := rows.Scan(&i.FollowingUserID, &i.FollowedUserID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unfollowUser = `-- name: UnfollowUser :exec
DELETE FROM follows
WHERE followed_user_id = $1 AND following_user_id = $2
`

type UnfollowUserParams struct {
	FollowedUserID  int32 `json:"followed_user_id"`
	FollowingUserID int32 `json:"following_user_id"`
}

func (q *Queries) UnfollowUser(ctx context.Context, arg UnfollowUserParams) error {
	_, err := q.db.Exec(ctx, unfollowUser, arg.FollowedUserID, arg.FollowingUserID)
	return err
}
