-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE Category(
    id SERIAL  PRIMARY KEY,
    name VARCHAR(100),
    created_at DATE,
    updated_at DATE
)

-- +migrate StatementEnd