CREATE TABLE "jobs" (
  "id" bigserial PRIMARY KEY,
  "description" text NOT NULL,
  "client_name" varchar NOT NULL,
  "client_address" varchar NOT NULL,
  "client_email" varchar NOT NULL,
  "price" bigint NOT NULL,
  "start_date" timestamptz NOT NULL,
  "end_date" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);