DO $$ BEGIN
    CREATE TYPE "public"."shop_status" AS ENUM (
        'ACTIVE'
        , 'INACTIVE'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS "shop" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    "owner_id" uuid NOT NULL,
    "name" varchar NOT NULL,
    "status" shop_status NOT NULL,

    "created_at" int8 NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())*1000,
    "created_by" varchar NOT NULL DEFAULT '',
    "updated_at" int8,
    "updated_by" varchar,
    "deleted_at" int8
);
