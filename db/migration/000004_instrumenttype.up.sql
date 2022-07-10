CREATE TABLE "instrumenttypes" (
  "id" bigserial PRIMARY KEY,
  "type" varchar NOT NULL,
  "alias" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL  DEFAULT (now())
);