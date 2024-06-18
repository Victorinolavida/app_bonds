CREATE TABLE IF NOT EXISTS bond_transaction (
    id bigserial PRIMARY KEY,
    transaction_id bigint not null,
    bond_id UUID not null
);