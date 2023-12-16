CREATE TABLE IF NOT EXISTS "user_account" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" varchar(128) UNIQUE NOT NULL,
    "password_hash" char (60) NOT NULL,
    "email" citext UNIQUE NOT NULL,
    "confirmation_token" varchar(100),
    "token_generation_time" timestamptz,
    "email_validation_status_id" int,
    "password_recovery_token" varchar(100),
    "password_recovery_expire" timestamptz,
    "user_status_id" int,
    "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX ON "user_account" ("username");

CREATE INDEX ON "user_account" ("email");