CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS book_category (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    category_name VARCHAR (250) 
);

CREATE TABLE IF NOT EXISTS books (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    author_name VARCHAR(250) NOT NULL,
    category_id uuid 
    REFERENCES book_category(id) ON DELETE SET NULL
);

INSERT INTO books (id,title, author_name, category_id ) 
VALUES ('a47a7a32-38e7-4934-a4b0-76497a4ba273','Kecha va kunduz','Cho"pon','48a42ec8-332e-4b1d-83b5-38ee9369ba4c');