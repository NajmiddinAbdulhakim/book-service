CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS book_category (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    category_name VARCHAR (250) 
);

CREATE TABLE IF NOT EXISTS books (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    title VARCHAR(250) NOT NULL UNIQUE,
    author_name VARCHAR(250) NOT NULL,
    category_id uuid 
    REFERENCES book_category(id) ON DELETE SET NULL
);



INSERT INTO book_category (category_name)
VALUES('Diniy'),('Roman'),('Qissa'),('Hikoyalar'),('Ilmiy-ommabop');