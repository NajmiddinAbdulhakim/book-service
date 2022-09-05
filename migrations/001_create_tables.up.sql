
CREATE TABLE IF NOT EXISTS book_category (
    id INTEGER PRIMARY KEY,
    category_name VARCHAR (250) 
);

CREATE TABLE IF NOT EXISTS books (
    id uuid NOT NULL PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    author_name VARCHAR(250) NOT NULL,
    category_id INTEGER 
    REFERENCES book_category(id) ON DELETE SET NULL
);