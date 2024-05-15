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

const GetTicketsByEventSlotId = `-- name: GetTicketsByEventSlotId :many
SELECT id, code, event_slot_id, status, price, issued, buyer_id FROM "user"."tickets"
WHERE event_slot_id = $1 
ORDER BY issued
`

func (q *Queries) GetTicketsByEventSlotId(ctx context.Context, eventSlotID uuid.UUID) ([]UserTicket, error) {
	rows, err := q.db.Query(ctx, GetTicketsByEventSlotId, eventSlotID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserTicket
	for rows.Next() {
		var i UserTicket
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.EventSlotID,
			&i.Status,
			&i.Price,
			&i.Issued,
			&i.BuyerID,
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

const GetTicketsByUserId = `-- name: GetTicketsByUserId :many
SELECT id, code, event_slot_id, status, price, issued, buyer_id FROM "user"."tickets"
WHERE buyer_id = $1
ORDER BY issued
`

func (q *Queries) GetTicketsByUserId(ctx context.Context, buyerID pgtype.UUID) ([]UserTicket, error) {
	rows, err := q.db.Query(ctx, GetTicketsByUserId, buyerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserTicket
	for rows.Next() {
		var i UserTicket
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.EventSlotID,
			&i.Status,
			&i.Price,
			&i.Issued,
			&i.BuyerID,
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

const GetTotalTicketByEventSlot = `-- name: GetTotalTicketByEventSlot :one
SELECT COUNT(*) FROM "user"."tickets"
WHERE event_slot_id = $1
`

func (q *Queries) GetTotalTicketByEventSlot(ctx context.Context, eventSlotID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, GetTotalTicketByEventSlot, eventSlotID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const SellTicket = `-- name: SellTicket :one
UPDATE "user".tickets as t
SET buyer_id = $1, STATUS = 'Booked'
WHERE ( id, code, $2) in (
	SELECT id , code, event_slot_id FROM "user".tickets
	WHERE buyer_id IS NULL AND event_slot_id = $2
	ORDER BY issued
	LIMIT 1
)  
RETURNING id, code, event_slot_id, status, price, issued, buyer_id
`

type SellTicketParams struct {
	BuyerID pgtype.UUID `json:"buyer_id"`
	Column2 interface{} `json:"column_2"`
}

func (q *Queries) SellTicket(ctx context.Context, arg SellTicketParams) (UserTicket, error) {
	row := q.db.QueryRow(ctx, SellTicket, arg.BuyerID, arg.Column2)
	var i UserTicket
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.EventSlotID,
		&i.Status,
		&i.Price,
		&i.Issued,
		&i.BuyerID,
	)
	return i, err
}
