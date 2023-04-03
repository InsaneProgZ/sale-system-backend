CREATE TABLE IF NOT EXISTS products (
    code int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    buy_value int,
    sell_value int NOT NULL,
    brand varchar(255) NOT NULL,
    creation_date timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (code)
);