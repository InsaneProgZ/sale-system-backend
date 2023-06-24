CREATE TABLE IF NOT EXISTS products (
    code int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL UNIQUE,
    buy_price int,
    sell_price int NOT NULL,
    brand varchar(255) NOT NULL,
    creation_date timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (code)
);