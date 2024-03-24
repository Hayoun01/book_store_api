-- name: CreateBook :one
INSERT INTO books (id, name, created_at, updated_at, published_at, isbn, description, author_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetAllBooks :many
SELECT * FROM books
ORDER BY created_at DESC
LIMIT $1;

-- name: GetBookByID :one
SELECT * FROM books
WHERE id = $1;

-- name: GetBookByIsbn :one
SELECT * FROM books
WHERE isbn = $1;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1;