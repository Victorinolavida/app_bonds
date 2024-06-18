ALTER TABLE bonds ADD CONSTRAINT name_max_length CHECK (LENGTH(name) <= 40 );
ALTER TABLE bonds ADD CONSTRAINT name_min_length CHECK (LENGTH(name) >= 4 );

ALTER TABLE bonds ADD CONSTRAINT bond_max CHECK (number_bonds <= 10000 );
ALTER TABLE bonds ADD CONSTRAINT bond_min CHECK (number_bonds >= 1 );

ALTER TABLE bonds ADD CONSTRAINT price_min CHECK (price >= 0);
ALTER TABLE bonds ADD CONSTRAINT price_max CHECK (price <=10000000000000 );

ALTER TABLE bonds ADD CONSTRAINT fk_owner_id FOREIGN KEY (owner_id) REFERENCES users(id);
ALTER TABLE bonds ADD CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id);