-- migrate:up
CREATE TABLE IF NOT EXISTS users (
    "id" UUID NOT NULL,
    "username" VARCHAR NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    PRIMARY KEY ("id")
)

-- migrate:down
drop table users