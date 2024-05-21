-- name: GetSource :one
SELECT * FROM sources
WHERE id = $1 
LIMIT 1;

-- name: CreateSource :one
INSERT INTO sources (
  name,
  data
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateSource :one
UPDATE sources 
SET data = $2, 
  name = $3
WHERE id = $1
RETURNING *;

-- name: GetSources :many
SELECT * FROM sources;