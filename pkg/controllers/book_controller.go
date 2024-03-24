package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	db "github.com/Hayoun01/book_store_api/pkg/db/sqlc"
	"github.com/Hayoun01/book_store_api/pkg/models"
	"github.com/Hayoun01/book_store_api/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func HandleGetReady(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJson(w, http.StatusOK, struct{}{})
}

func HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name        string    `json:"name"`
		AuthorId    uuid.UUID `json:"author_id"`
		Description string    `json:"description"`
		PublishedAt string    `json:"published_at"`
		Isbn        string    `json:"isbn"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.ResponseWithJson(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	layout := "02-01-2006"
	parsedDate, err := time.Parse(layout, params.PublishedAt)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Invalid date format: %v, Excepted format: %v", params.PublishedAt, layout))
		return
	}
	isbn := sql.NullString{}
	if params.Isbn != "" {
		isbn.String = string(params.Isbn)
		isbn.Valid = true
	}
	description := sql.NullString{}
	if params.Description != "" {
		description.String = string(params.Description)
		description.Valid = true
	}
	book, err := models.ApiCfg.DB.CreateBook(r.Context(), db.CreateBookParams{
		ID:          uuid.New(),
		Name:        params.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		PublishedAt: parsedDate,
		Isbn:        isbn,
		Description: description,
		AuthorID:    params.AuthorId,
	})
	if err != nil {
		if strings.Contains(err.Error(), "books_isbn_key") {
			utils.ResponseWithError(w, 400, "There is a book with the same ISBN already.")
			return
		}
		utils.ResponseWithError(w, 400, fmt.Sprintf("Error creating this book: %v", err.Error()))
		return
	}
	utils.ResponseWithJson(w, 201, models.BookToBook(book))
}

func HandleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.ApiCfg.DB.GetAllBooks(r.Context(), 10)
	if err != nil {
		utils.ResponseWithError(w, 400, err.Error())
		return
	}
	utils.ResponseWithJson(w, 200, models.BooksToBooks(books))
}

func HandleGetBookByID(w http.ResponseWriter, r *http.Request) {
	book_id := mux.Vars(r)["book_id"]
	if book_id == "" {
		utils.ResponseWithError(w, 400, "no book_id passed")
		return
	}
	parsedId, err := uuid.Parse(book_id)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Invalid book_id: %v", err.Error()))
		return
	}
	book, err := models.ApiCfg.DB.GetBookByID(r.Context(), parsedId)
	if err != nil {
		utils.ResponseWithError(w, 400, err.Error())
		return
	}
	utils.ResponseWithJson(w, 200, models.BookToBook(book))
}

func HandleDeleteBookByID(w http.ResponseWriter, r *http.Request) {
	book_id := mux.Vars(r)["book_id"]
	if book_id == "" {
		utils.ResponseWithError(w, 400, "No book_id passed")
		return
	}
	parsedID, err := uuid.Parse(book_id)
	if err != nil {
		utils.ResponseWithError(w, 200, "Invalid book_id format.")
		return
	}
	err = models.ApiCfg.DB.DeleteBook(r.Context(), parsedID)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("couldn't delete this book due: %v", err.Error()))
		return
	}
	utils.ResponseWithJson(w, 200, struct{}{})
}
