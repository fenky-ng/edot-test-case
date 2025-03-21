CREATE TABLE IF NOT EXISTS "user" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    "name" varchar NOT NULL,
    "phone" varchar,
    "email" varchar,
    "hashed_password" varchar NOT NULL,

    "created_at" int8 NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())*1000,
    "created_by" varchar NOT NULL DEFAULT '',
    "updated_at" int8,
    "updated_by" varchar,
    "deleted_at" int8
);
