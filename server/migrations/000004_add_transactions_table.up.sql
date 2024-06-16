CREATE TABLE IF NOT EXISTS transactions (
    id bigserial  PRIMARY KEY,
    bond_id bigint not null,
    seller_id bigint not null,
    buyer_id bigint not null,
    deleted_at timestamp(0) with time zone DEFAULT NULL
);