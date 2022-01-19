CREATE EXTENSION IF not exists "uuid-ossp";

CREATE TABLE IF NOT EXISTS contacts (
    "id" uuid DEFAULT uuid_generate_v4 (),
    "firstname" varchar NOT NULL,
    "lastname" varchar NOT NULL,
    "email" varchar NOT NULL,
    "phone" varchar NOT NULL,
    "position" varchar NOT NULL,
    PRIMARY KEY ("id")

);

CREATE TABLE IF NOT EXISTS task (
    "id" uuid DEFAULT uuid_generate_v4 (),
    "name" varchar NOT NULL,
    "status" varchar NOT NULL,
    "priority" varchar NOT NULL,
    "created_at" varchar NOT NULL,
    "created_by" uuid,
    "due_date" varchar NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY(created_by) references contacts(id) on delete cascade
);
