-- name: GetEvent :one
SELECT * FROM "events"
WHERE id = $1 
LIMIT 1;

-- name: CreateEvent :one
INSERT INTO "events" (
  event_name,
  note,
  event_owner_id  
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: StartEvent :one
UPDATE "events"
SET STATUS = 'Open', updated = now()
WHERE id = $1
RETURNING *;

-- name: OpenEventSlots :one
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
RETURNING *;

-- name: UpdateEventSlots :one
UPDATE event_slots 
SET slot_name = $2
WHERE id = $1
RETURNING *;

-- name: CloseEventSlot :one
UPDATE event_slots 
SET status = 'Closed'
WHERE id = $1
RETURNING *;


-- name: GetEventSlotsByEventId :many
SELECT * FROM "event_slots"
WHERE event_id = $1;

-- name: GetEventSlotsById :one
SELECT * FROM "event_slots"
WHERE id = $1;
