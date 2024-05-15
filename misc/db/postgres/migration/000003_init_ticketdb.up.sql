CREATE SCHEMA "user";



CREATE TYPE ticket_status AS ENUM ('New', 'Booked', 'CheckedIn' , 'InStock');

CREATE TABLE "user"."tickets" (
  "id" uuid NOT NULL DEFAULT (uuid_generate_v4()),
  "code" varchar(100) NOT NULL,
  "event_slot_id" uuid NOT NULL,
  "status" ticket_status NOT NULL DEFAULT ('New'),
  "price" real NOT NULL,
  "issued" timestamp NOT NULL DEFAULT (now()),
  "buyer_id" uuid NULL,
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "ix_user_tickets_id" ON "user"."tickets" ("id");

CREATE INDEX "ix_user_tickets_event_slot_id" ON "user"."tickets" ("event_slot_id");
