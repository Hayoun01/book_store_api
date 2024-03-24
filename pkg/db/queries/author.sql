-- name: CreateAuthor :one
INSERT INTO authors (id, name, created_at, updated_at, description) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAllAuthors :many
SELECT * FROM authors
ORDER BY name ASC
LIMIT $1;