CREATE TABLE IF NOT EXISTS users (
    id           bigserial PRIMARY KEY,
    createdAt    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updatedAt    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    username     text                        NOT NULL,
    password     text                        NOT NULL
);