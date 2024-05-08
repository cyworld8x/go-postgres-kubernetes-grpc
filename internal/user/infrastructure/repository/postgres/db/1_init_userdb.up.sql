CREATE SCHEMA "db" ;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "db"."users" (
  "id" uuid NOT NULL DEFAULT (uuid_generate_v4()),
  "code" varchar(100) NOT NULL,
  "username" varchar(100) NOT NULL,
  "display_name" varchar(100),
  "password" varchar NOT NULL,
  "email" varchar,
  "status" boolean NOT NULL DEFAULT true,
  "role" integer NOT NULL,
  "created" timestamp NOT NULL DEFAULT (now()),
  "updated" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "ix_user_users_id" ON "db"."users" ("id");