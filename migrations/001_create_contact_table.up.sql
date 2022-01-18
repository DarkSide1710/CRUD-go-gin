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

CREATE TABLE IF NOT EXISTS tasks (
    "id" uuid DEFAULT uuid_generate_v4 (),
    "name" varchar NOT NULL,
    "status" varchar NOT NULL,
    "priority" varchar NOT NULL,
    "createdat" varchar NOT NULL,
    "createby" varchar,
    "duedate" varchar NOT NULL,
     FOREIGN KEY(createby) references contacts(id)
);

