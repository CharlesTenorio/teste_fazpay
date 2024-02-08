CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tb_product
(
    id serial NOT NULL,
    external_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name character varying(120) NOT NULL,
    quantity INTEGER NOT NULL,
    price double precision NOT NULL,
    created_at time without time zone NOT NULL DEFAULT NOW(),
    updated_at time without time zone NOT NULL DEFAULT NOW(),
    CONSTRAINT tb_product_pkey PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON tb_product
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

INSERT INTO tb_product (name, quantity, price) VALUES 
    ('Notebook 13 XPS', 10, 14500),
    ('Notebook 15', 15, 2500),
    ('Notebook 14', 20, 2500),
    ('Notebook 15', 12, 3500),
    ('Notebook 13', 8, 4568),
    ('Tablet', 5, 8450),
    ('Macbook 13 Pro M1', 30, 18500.00),
    ('TV 55', 25, 4500.00),
    ('TV 45', 18, 3500.00),
    ('TV 32', 22, 2500.00),
    ('TV 60', 10, 6500.00),
    ('TV 50', 15, 4800.00);
