CREATE TYPE ticket_status AS ENUM ('OPEN', 'CLOSED');

CREATE TYPE vehicle_type AS ENUM ('Car', 'Bike', 'Truck');

CREATE TABLE "ticket" (
  "id" varchar PRIMARY KEY,
  "vehicle_colour" varchar NOT NULL,
  "vehicle_type" vehicle_type NOT NULL,
  "registration_no" varchar NOT NULL,
  "status" ticket_status NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "slot" (
  "id" integer PRIMARY KEY,
  "floor_id" integer NOT NULL,
  "vehicle_type" vehicle_type NOT NULL,
  "is_booked" boolean NOT NULL
);

CREATE TABLE "floor" (
  "id" integer PRIMARY KEY,
  "parking_lot_id" varchar NOT NULL
);

CREATE INDEX ON "ticket" ("id");

CREATE INDEX ON "slot" ("vehicle_type");