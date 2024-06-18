CREATE TABLE IF NOT EXISTS transactions (
    id bigserial  PRIMARY KEY,
    seller_id bigint not null,
    price bigint not null,
    buyer_id bigint not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);