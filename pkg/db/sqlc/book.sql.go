// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: book.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (id, name, created_at, updated_at, published_at, isbn, description, author_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, created_at, updated_at, published_at, isbn, description, author_id
`

type CreateBookParams struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	PublishedAt time.Time      `json:"published_at"`
	Isbn        sql.NullString `json:"isbn"`
	Description sql.NullString `json:"description"`
	AuthorID    uuid.UUID      `json:"author_id"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.PublishedAt,
		arg.Isbn,
		arg.Description,
		arg.AuthorID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PublishedAt,
		&i.Isbn,
		&i.Description,
		&i.AuthorID,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const getAllBooks = `-- name: GetAllBooks :many
SELECT id, name, created_at, updated_at, published_at, isbn, description, author_id FROM books
ORDER BY created_at DESC
LIMIT $1
`

func (q *Queries) GetAllBooks(ctx context.Context, limit int32) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, getAllBooks, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PublishedAt,
			&i.Isbn,
			&i.Description,
			&i.AuthorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookByID = `-- name: GetBookByID :one
SELECT id, name, created_at, updated_at, published_at, isbn, description, author_id FROM books
WHERE id = $1
`

func (q *Queries) GetBookByID(ctx context.Context, id uuid.UUID) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBookByID, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PublishedAt,
		&i.Isbn,
		&i.Description,
		&i.AuthorID,
	)
	return i, err
}

const getBookByIsbn = `-- name: GetBookByIsbn :one
SELECT id, name, created_at, updated_at, published_at, isbn, description, author_id FROM books
WHERE isbn = $1
`

func (q *Queries) GetBookByIsbn(ctx context.Context, isbn sql.NullString) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBookByIsbn, isbn)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PublishedAt,
		&i.Isbn,
		&i.Description,
		&i.AuthorID,
	)
	return i, err
}