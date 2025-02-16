CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    final_price FLOAT NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO
    orders (id, price, tax, final_price)
VALUES
    ('abababa', 100.5, 0.5, 101),
    ('cdcdcdc', 200.0, 10.0, 210),
    ('efefefe', 50.75, 2.5, 53.25),
    ('ghghghg', 300.99, 15.0, 315.99),
    ('ijkijki', 120.45, 5.55, 126.00),
    ('lmlmlml', 89.90, 4.10, 94.00) ON DUPLICATE KEY
UPDATE
    id =
VALUES
(id);