-- +goose Up
CREATE TABLE
    authors (
        id UUID PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        description text
    );

-- +goose Down
DROP TABLE authors;