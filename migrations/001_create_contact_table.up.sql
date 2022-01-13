CREATE TABLE IF NOT EXISTS "contacts" (
    "id" int NOT NULL PRIMARY KEY,
    "firstname" varchar NOT NULL,
    "lastname" varchar NOT NULL,
    "email" varchar NOT NULL,
    "phone" varchar NOT NULL,
    "position" varchar NOT NULL
);