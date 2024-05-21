CREATE TABLE Sources (
  "id" uuid NOT NULL DEFAULT (uuid_generate_v4()),
  "name" varchar(100) NOT NULL,
  "data" json    NOT NULL,
  "created" timestamp NOT NULL DEFAULT (now()),
  "updated" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "ix_sources_id" ON "sources" ("id");