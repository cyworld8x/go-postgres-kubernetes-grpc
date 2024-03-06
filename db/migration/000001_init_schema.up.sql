CREATE TABLE "follows" (
  "following_user_id" integer NOT NULL,
  "followed_user_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "fullname" varchar,
  "password" varchar,
  "role" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "body" text,
  "user_id" integer,
  "status" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("following_user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id");
