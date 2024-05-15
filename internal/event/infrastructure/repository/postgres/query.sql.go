// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const CloseEventSlot = `-- name: CloseEventSlot :one
UPDATE event_slots 
SET status = 'Closed'
WHERE id = $1
RETURNING id, slot_name, description, price, capacity, status, start_time, end_time, event_id, created, updated
`

func (q *Queries) CloseEventSlot(ctx context.Context, id uuid.UUID) (EventSlot, error) {
	row := q.db.QueryRow(ctx, CloseEventSlot, id)
	var i EventSlot
	err := row.Scan(
		&i.ID,
		&i.SlotName,
		&i.Description,
		&i.Price,
		&i.Capacity,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.EventID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const CreateEvent = `-- name: CreateEvent :one
INSERT INTO "events" (
  event_name,
  note,
  event_owner_id  
) VALUES (
  $1, $2, $3
)
RETURNING id, event_name, note, revenue, status, total_sold_tickets, event_owner_id, created, updated
`

type CreateEventParams struct {
	EventName    string    `json:"event_name"`
	Note         string    `json:"note"`
	EventOwnerID uuid.UUID `json:"event_owner_id"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, CreateEvent, arg.EventName, arg.Note, arg.EventOwnerID)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.EventName,
		&i.Note,
		&i.Revenue,
		&i.Status,
		&i.TotalSoldTickets,
		&i.EventOwnerID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const GetEvent = `-- name: GetEvent :one
SELECT id, event_name, note, revenue, status, total_sold_tickets, event_owner_id, created, updated FROM "events"
WHERE id = $1 
LIMIT 1
`

func (q *Queries) GetEvent(ctx context.Context, id uuid.UUID) (Event, error) {
	row := q.db.QueryRow(ctx, GetEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.EventName,
		&i.Note,
		&i.Revenue,
		&i.Status,
		&i.TotalSoldTickets,
		&i.EventOwnerID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const GetEventSlotsByEventId = `-- name: GetEventSlotsByEventId :many
SELECT id, slot_name, description, price, capacity, status, start_time, end_time, event_id, created, updated FROM "event_slots"
WHERE event_id = $1
`

func (q *Queries) GetEventSlotsByEventId(ctx context.Context, eventID uuid.UUID) ([]EventSlot, error) {
	rows, err := q.db.Query(ctx, GetEventSlotsByEventId, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EventSlot
	for rows.Next() {
		var i EventSlot
		if err := rows.Scan(
			&i.ID,
			&i.SlotName,
			&i.Description,
			&i.Price,
			&i.Capacity,
			&i.Status,
			&i.StartTime,
			&i.EndTime,
			&i.EventID,
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

const GetEventSlotsById = `-- name: GetEventSlotsById :one
SELECT id, slot_name, description, price, capacity, status, start_time, end_time, event_id, created, updated FROM "event_slots"
WHERE id = $1
`

func (q *Queries) GetEventSlotsById(ctx context.Context, id uuid.UUID) (EventSlot, error) {
	row := q.db.QueryRow(ctx, GetEventSlotsById, id)
	var i EventSlot
	err := row.Scan(
		&i.ID,
		&i.SlotName,
		&i.Description,
		&i.Price,
		&i.Capacity,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.EventID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const OpenEventSlots = `-- name: OpenEventSlots :one
INSERT INTO "event_slots" (
  slot_name,
  description,
  price,
  capacity,
  start_time,
  end_time,
  event_id  
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, slot_name, description, price, capacity, status, start_time, end_time, event_id, created, updated
`

type OpenEventSlotsParams struct {
	SlotName    string           `json:"slot_name"`
	Description string           `json:"description"`
	Price       float32          `json:"price"`
	Capacity    int32            `json:"capacity"`
	StartTime   pgtype.Timestamp `json:"start_time"`
	EndTime     pgtype.Timestamp `json:"end_time"`
	EventID     uuid.UUID        `json:"event_id"`
}

func (q *Queries) OpenEventSlots(ctx context.Context, arg OpenEventSlotsParams) (EventSlot, error) {
	row := q.db.QueryRow(ctx, OpenEventSlots,
		arg.SlotName,
		arg.Description,
		arg.Price,
		arg.Capacity,
		arg.StartTime,
		arg.EndTime,
		arg.EventID,
	)
	var i EventSlot
	err := row.Scan(
		&i.ID,
		&i.SlotName,
		&i.Description,
		&i.Price,
		&i.Capacity,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.EventID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const StartEvent = `-- name: StartEvent :one
UPDATE "events"
SET STATUS = 'Open', updated = now()
WHERE id = $1
RETURNING id, event_name, note, revenue, status, total_sold_tickets, event_owner_id, created, updated
`

func (q *Queries) StartEvent(ctx context.Context, id uuid.UUID) (Event, error) {
	row := q.db.QueryRow(ctx, StartEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.EventName,
		&i.Note,
		&i.Revenue,
		&i.Status,
		&i.TotalSoldTickets,
		&i.EventOwnerID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const UpdateEventSlots = `-- name: UpdateEventSlots :one
UPDATE event_slots 
SET slot_name = $2
WHERE id = $1
RETURNING id, slot_name, description, price, capacity, status, start_time, end_time, event_id, created, updated
`

type UpdateEventSlotsParams struct {
	ID       uuid.UUID `json:"id"`
	SlotName string    `json:"slot_name"`
}

func (q *Queries) UpdateEventSlots(ctx context.Context, arg UpdateEventSlotsParams) (EventSlot, error) {
	row := q.db.QueryRow(ctx, UpdateEventSlots, arg.ID, arg.SlotName)
	var i EventSlot
	err := row.Scan(
		&i.ID,
		&i.SlotName,
		&i.Description,
		&i.Price,
		&i.Capacity,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.EventID,
		&i.Created,
		&i.Updated,
	)
	return i, err
}
