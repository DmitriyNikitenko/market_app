CREATE TABLE IF NOT EXISTS products (
    id           SERIAL         PRIMARY KEY,
    name         VARCHAR(100)   NOT NULL,
    price        NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    category     VARCHAR(50)    NOT NULL,
    is_available BOOLEAN        DEFAULT TRUE
);

