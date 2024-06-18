CREATE TABLE IF NOT EXISTS users (
    id bigserial not null PRIMARY KEY,
    email  TEXT not null UNIQUE,
    username TEXT  NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    password_hash bytea NOT NULL
);
