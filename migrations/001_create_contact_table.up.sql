CREATE TABLE IF NOT EXISTS "contacts" (
    "id" varchar NOT NULL PRIMARY KEY,
    "firstname" varchar NOT NULL,
    "lastname" varchar NOT NULL,
    "email" varchar NOT NULL,
    "phone" varchar NOT NULL,
    "position" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "task" (
    "id" int NOT NULL PRIMARY KEY,
    "name" varchar NOT NULL,
    "status" varchar NOT NULL,
    "priority" varchar NOT NULL,
    "createdat" varchar NOT NULL,
    "createby" varchar references contacts(id),
    "duedate" varchar NOT NULL
);
