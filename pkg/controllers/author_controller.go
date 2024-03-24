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
)

func HandleCreateAuthor(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.ResponseWithJson(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	description := sql.NullString{}
	if params.Description != "" {
		description.String = params.Description
		description.Valid = true
	}
	author, err := models.ApiCfg.DB.CreateAuthor(r.Context(), db.CreateAuthorParams{
		ID:          uuid.New(),
		Name:        params.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: description,
	})
	if err != nil {
		if strings.Contains(err.Error(), "books_isbn_key") {
			utils.ResponseWithError(w, 400, "There is a book with the same ISBN already.")
			return
		}
		utils.ResponseWithError(w, 400, fmt.Sprintf("Error creating this book: %v", err.Error()))
		return
	}
	utils.ResponseWithJson(w, 201, models.AuthorToAuthor(author))
}

func HandleGetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := models.ApiCfg.DB.GetAllAuthors(r.Context(), 10)
	if err != nil {
		utils.ResponseWithError(w, 400, err.Error())
		return
	}
	utils.ResponseWithJson(w, 200, models.AuthorsToAuthors(authors))
}
