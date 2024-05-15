-- name: GetTicketsByEventSlotId :many
SELECT * FROM "user"."tickets"
WHERE event_slot_id = $1 
ORDER BY issued;

-- name: GetTicketsByUserId :many
SELECT * FROM "user"."tickets"
WHERE buyer_id = $1
ORDER BY issued;

-- name: SellTicket :one
UPDATE "user".tickets as t
SET buyer_id = $1, STATUS = 'Booked'
WHERE ( id, code, $2) in (
	SELECT id , code, event_slot_id FROM "user".tickets
	WHERE buyer_id IS NULL AND event_slot_id = $2
	ORDER BY issued
	LIMIT 1
)  
RETURNING *;

-- name: GetTotalTicketByEventSlot :one
SELECT COUNT(*) FROM "user"."tickets"
WHERE event_slot_id = $1;
