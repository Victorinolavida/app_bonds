ALTER TABLE bonds ADD CONSTRAINT name_max_length CHECK (LENGTH(name) <= 40 );
ALTER TABLE bonds ADD CONSTRAINT name_min_length CHECK (LENGTH(name) >= 4 );

ALTER TABLE bonds ADD CONSTRAINT bond_max CHECK (number_bonds <= 10000 );
ALTER TABLE bonds ADD CONSTRAINT bond_min CHECK (number_bonds >= 1 );

ALTER TABLE bonds ADD CONSTRAINT price_min CHECK (price >= 0.0000 );
ALTER TABLE bonds ADD CONSTRAINT price_max CHECK (price <= 100000000.0000 );

ALTER TABLE bonds ADD CONSTRAINT foreign_key_issuer FOREIGN KEY (issuer_id) 
REFERENCES users(id) ON DELETE CASCADE;