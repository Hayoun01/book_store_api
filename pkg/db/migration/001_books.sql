-- +goose Up
CREATE TABLE
    books (
        id UUID PRIMARY KEY,
        name TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        published_at DATE NOT NULL,
        isbn VARCHAR(20) UNIQUE
    );

-- +goose Down
DROP TABLE books;