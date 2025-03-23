DO $$ BEGIN
    CREATE TYPE "public"."warehouse_status" AS ENUM (
        'ACTIVE'
        , 'INACTIVE'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS "warehouse" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    "shop_id" uuid NOT NULL,
    "name" varchar NOT NULL,
    "status" warehouse_status NOT NULL,

    "created_at" int8 NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())*1000,
    "created_by" varchar NOT NULL DEFAULT '',
    "updated_at" int8,
    "updated_by" varchar,
    "deleted_at" int8
);

CREATE TABLE IF NOT EXISTS "stock" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    "warehouse_id" uuid NOT NULL,
    "product_id" varchar NOT NULL,
    "stock" int8 NOT NULL,

    "created_at" int8 NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())*1000,
    "created_by" varchar NOT NULL DEFAULT '',
    "updated_at" int8,
    "updated_by" varchar,
    "deleted_at" int8,

    CONSTRAINT stock_unique_warehouse_id_product_id UNIQUE (warehouse_id, product_id)
);
