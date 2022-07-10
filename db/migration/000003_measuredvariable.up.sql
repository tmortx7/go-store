CREATE TABLE "measuredvariables" (
  "id" bigserial PRIMARY KEY,
  "variable" varchar NOT NULL,
  "alias" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL  DEFAULT (now())
)