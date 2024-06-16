
ALTER TABLE bonds DROP CONSTRAINT  name_max_length;
ALTER TABLE bonds DROP CONSTRAINT  name_min_length;

ALTER TABLE bonds DROP CONSTRAINT  bond_max;
ALTER TABLE bonds DROP CONSTRAINT  bond_min;

ALTER TABLE bonds DROP CONSTRAINT price_min;
ALTER TABLE bonds DROP CONSTRAINT price_max;

ALTER TABLE bonds DROP CONSTRAINT foreign_key_issuer;
