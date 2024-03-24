package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/Hayoun01/book_store_api/pkg/config"
	db "github.com/Hayoun01/book_store_api/pkg/db/sqlc"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

var ApiCfg = apiConfig{}

func init() {
	err := godotenv.Load("./pkg/config/.env")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open("postgres", config.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	database := db.New(conn)
	ApiCfg.DB = database
}

type Book struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
	Isbn        string    `json:"isbn"`
	Description string    `json:"description"`
	AuthorID    uuid.UUID `json:"author_id"`
}

func BookToBook(dbBook db.Book) Book {
	return Book{
		ID:          dbBook.ID,
		Name:        dbBook.Name,
		CreatedAt:   dbBook.CreatedAt,
		UpdatedAt:   dbBook.UpdatedAt,
		PublishedAt: dbBook.PublishedAt,
		Isbn:        dbBook.Isbn.String,
		Description: dbBook.Description.String,
		AuthorID:    dbBook.AuthorID,
	}
}

func BooksToBooks(dbBook []db.Book) []Book {
	books := []Book{}
	for _, book := range dbBook {
		books = append(books, BookToBook(book))
	}
	return books
}

type Author struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
}

func AuthorToAuthor(dbBook db.Author) Author {
	return Author{
		ID:          dbBook.ID,
		Name:        dbBook.Name,
		CreatedAt:   dbBook.CreatedAt,
		UpdatedAt:   dbBook.UpdatedAt,
		Description: dbBook.Description.String,
	}
}

func AuthorsToAuthors(dbBook []db.Author) []Author {
	authors := []Author{}
	for _, author := range dbBook {
		authors = append(authors, AuthorToAuthor(author))
	}
	return authors
}
