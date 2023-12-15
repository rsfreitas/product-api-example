DROP TABLE IF EXISTS product;
CREATE TABLE product (
    id varchar(255) not null,
    name varchar(255) not null,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id)
);
