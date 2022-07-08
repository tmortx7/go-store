CREATE TABLE "processvariables" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "alias" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);