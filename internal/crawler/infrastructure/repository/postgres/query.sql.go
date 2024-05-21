// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
)

const CreateSource = `-- name: CreateSource :one
INSERT INTO sources (
  name,
  data
) VALUES (
  $1, $2
)
RETURNING id, name, data, created, updated
`

type CreateSourceParams struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

func (q *Queries) CreateSource(ctx context.Context, arg CreateSourceParams) (Source, error) {
	row := q.db.QueryRow(ctx, CreateSource, arg.Name, arg.Data)
	var i Source
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Data,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const GetSource = `-- name: GetSource :one
SELECT id, name, data, created, updated FROM sources
WHERE id = $1 
LIMIT 1
`

func (q *Queries) GetSource(ctx context.Context, id uuid.UUID) (Source, error) {
	row := q.db.QueryRow(ctx, GetSource, id)
	var i Source
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Data,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const GetSources = `-- name: GetSources :many
SELECT id, name, data, created, updated FROM sources
`

func (q *Queries) GetSources(ctx context.Context) ([]Source, error) {
	rows, err := q.db.Query(ctx, GetSources)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Source
	for rows.Next() {
		var i Source
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Data,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateSource = `-- name: UpdateSource :one
UPDATE sources 
SET data = $2, 
  name = $3
WHERE id = $1
RETURNING id, name, data, created, updated
`

type UpdateSourceParams struct {
	ID   uuid.UUID `json:"id"`
	Data []byte    `json:"data"`
	Name string    `json:"name"`
}

func (q *Queries) UpdateSource(ctx context.Context, arg UpdateSourceParams) (Source, error) {
	row := q.db.QueryRow(ctx, UpdateSource, arg.ID, arg.Data, arg.Name)
	var i Source
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Data,
		&i.Created,
		&i.Updated,
	)
	return i, err
}
