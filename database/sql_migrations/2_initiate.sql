-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE Books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(250) NOT NULL,
    image_url VARCHAR(250) NOT NULL,
    release_year INT NOT NULL,
    price VARCHAR(50) NOT NULL,
    total_page INT NOT NULL,
    thickness VARCHAR (50),
    created_at DATE,
    updated_at DATE,
    category_id INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES Category(id)
)

-- +migrate StatementEnd