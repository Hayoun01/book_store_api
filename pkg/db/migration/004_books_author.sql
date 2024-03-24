-- +goose Up
ALTER TABLE books
ADD COLUMN author_id uuid NOT NULL REFERENCES authors (id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE books
DROP COLUMN author_id;