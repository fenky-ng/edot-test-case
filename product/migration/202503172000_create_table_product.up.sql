DO $$ BEGIN
    CREATE TYPE "public"."product_status" AS ENUM (
        'ACTIVE'
        , 'INACTIVE'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS "product" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    "shop_id" uuid NOT NULL,
    "name" varchar NOT NULL,
    "description" text NOT NULL,
    "price" int8 NOT NULL,
    "status" product_status NOT NULL,

    "created_at" int8 NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())*1000,
    "created_by" varchar NOT NULL DEFAULT '',
    "updated_at" int8,
    "updated_by" varchar,
    "deleted_at" int8
);
