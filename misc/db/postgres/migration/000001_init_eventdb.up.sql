
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "events" (
  "id" uuid NOT NULL DEFAULT (uuid_generate_v4()),
  "event_name" varchar(200) NOT NULL,
  "event_owner_id" uuid NOT NULL,
  "created" timestamp NOT NULL DEFAULT (now()),
  "updated" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "event_slots" (
  "id" uuid NOT NULL DEFAULT (uuid_generate_v4()),
  "item_type" integer NOT NULL,
  "slot_name" varchar(200) NOT NULL,
  "description" text NOT NULL,
  "price" money NOT NULL,
  "capacity" integer NOT NULL,
  "status" integer NOT NULL,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp NOT NULL,
  "event_id" uuid NOT NULL,
  "created" timestamp NOT NULL DEFAULT (now()),
  "updated" timestamp,
  PRIMARY KEY ("id")
);


CREATE UNIQUE INDEX "ix_events_id" ON "events" ("id");

CREATE UNIQUE INDEX "ix_event_slots_id" ON "event_slots" ("id");

CREATE INDEX "ix_event_slots_event_id" ON "event_slots" ("event_id");

ALTER TABLE "event_slots" ADD CONSTRAINT "fk_event_slots_events_event_id" FOREIGN KEY ("event_id") REFERENCES "events" ("id");
