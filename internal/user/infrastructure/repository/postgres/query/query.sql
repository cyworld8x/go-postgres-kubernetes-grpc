-- name: GetUser :one
SELECT * FROM "db"."users"
WHERE id = $1 
LIMIT 1;

-- name: GetLogin :one
SELECT * FROM "db"."users"
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "db"."users"
ORDER BY username;

-- name: CreateUser :one
INSERT INTO "db"."users" (
  username, email, display_name, password, role, code
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateUserInfo :exec
UPDATE "db"."users"
  set email = $4,
  display_name = $3
WHERE username = $2 AND id = $1;

-- name: ChangeUserPassword :exec
UPDATE "db"."users"
  set password = $3,
  updated = now()
WHERE username = $2 AND id = $1;

-- name: DeactiveUser :exec
UPDATE "db"."users"
  set status = false
WHERE id = $1;