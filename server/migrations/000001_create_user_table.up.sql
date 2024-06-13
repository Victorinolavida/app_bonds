CREATE TABLE IF NOT EXISTS users (
    id bigserial not null PRIMARY KEY,
    email  TEXT not null UNIQUE,
    username TEXT  NOT NULL UNIQUE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    password_hash bytea NOT NULL
);
