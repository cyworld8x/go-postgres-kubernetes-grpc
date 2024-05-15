
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE event_status AS ENUM ('Planned', 'Open', 'Closed', 'Cancelled');

CREATE TYPE event_slot_status AS ENUM ('New', 'Closed', 'Cancelled');


CREATE TABLE "events" (
  "id" uuid NOT NULL DEFAULT (uuid_generate_v4()),
  "event_name" varchar(200) NOT NULL,
  "note" varchar(200) NOT NULL,
  "revenue" real NOT NULL DEFAULT (0.0),
  "status" event_status NOT NULL DEFAULT 'Planned',
  "total_sold_tickets" integer NOT NULL DEFAULT 0,
  "event_owner_id" uuid NOT NULL,
  "created" timestamp NOT NULL DEFAULT (now()),
  "updated" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "event_slots" (
  "id" uuid NOT NULL DEFAULT (uuid_generate_v4()),
  "slot_name" varchar(200) NOT NULL,
  "description" text NOT NULL,
  "price" real NOT NULL,
  "capacity" integer NOT NULL,
  "status" event_slot_status NOT NULL DEFAULT 'New',
  "start_time" timestamp NOT NULL,
  "end_time" timestamp NOT NULL,
  "event_id" uuid NOT NULL,
  "created" timestamp NOT NULL DEFAULT (now()),
  "updated" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);


CREATE UNIQUE INDEX "ix_events_id" ON "events" ("id");

CREATE UNIQUE INDEX "ix_event_slots_id" ON "event_slots" ("id");

CREATE INDEX "ix_event_slots_event_id" ON "event_slots" ("event_id");

ALTER TABLE "event_slots" ADD CONSTRAINT "fk_event_slots_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id");
