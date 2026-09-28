CREATE TABLE IF NOT EXISTS products (
    id           SERIAL         PRIMARY KEY,
    name         VARCHAR(100)   NOT NULL,
    price        NUMERIC(10, 2) NOT NULL,
    category     VARCHAR(50)    NOT NULL,
    is_available BOOLEAN        DEFAULT TRUE
);

INSERT INTO products (name, price, category, is_available) VALUES
('Клубника в молочном шоколаде', 100.00, 'Клубника', true),
('Клубника в белом шоколаде', 100.00, 'Клубника', true),
('Латте', 150.00, 'Кофе', true);