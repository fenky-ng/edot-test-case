DO $$ BEGIN
    CREATE TYPE "public"."order_status" AS ENUM (
        'WAITING_FOR_PAYMENT'
        , 'CANCELLED_STOCK_ISSUE'
        , 'EXPIRED'
        , 'PAID'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS "order" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    "user_id" uuid NOT NULL,
    "order_no" varchar NOT NULL,
    "status" order_status NOT NULL,
    -- "recipient_name" varchar NOT NULL,
    -- "recipient_contact" varchar NOT NULL,
    -- "recipient_address" varchar NOT NULL,
    "payment_ref_no" varchar,
    "error_message" varchar,

    "created_at" int8 NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())*1000,
    "created_by" varchar NOT NULL DEFAULT '',
    "updated_at" int8,
    "updated_by" varchar,
    "deleted_at" int8
);

CREATE TABLE IF NOT EXISTS "order_detail" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    "order_id" uuid NOT NULL,
    "product_id" uuid NOT NULL,
    "warehouse_id" uuid NOT NULL,
    "price" int8 NOT NULL,
    "quantity" int8 NOT NULL,

    "created_at" int8 NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())*1000,
    "created_by" varchar NOT NULL DEFAULT '',
    "updated_at" int8,
    "updated_by" varchar,
    "deleted_at" int8
);
