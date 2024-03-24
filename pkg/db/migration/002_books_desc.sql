-- +goose Up
ALTER TABLE books
ADD COLUMN description text;

-- +goose Down
ALTER TABLE books
DROP COLUMN description;