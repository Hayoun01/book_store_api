package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/Hayoun01/book_store_api/pkg/db/sqlc"
	"github.com/Hayoun01/book_store_api/pkg/models"
	"github.com/Hayoun01/book_store_api/pkg/utils"
	"github.com/google/uuid"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := models.ApiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err.Error()))
		return
	}
	utils.ResponseWithJson(w, 201, user)
}
