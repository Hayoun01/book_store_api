package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Hayoun01/book_store_api/pkg/models"
	"github.com/Hayoun01/book_store_api/pkg/utils"
)

func GetApiKEY(r *http.Request) (string, error) {
	key := r.Header.Get("Authorization")
	if key == "" {
		return "", errors.New("no authentication key found")
	}
	slice := strings.Split(key, " ")
	if len(slice) != 2 {
		return "", errors.New("malformed authorization header")
	}
	if slice[0] != "ApiKey" {
		return "", errors.New("malformed API key")
	}
	return slice[1], nil
}

type authedHandler func(http.ResponseWriter, *http.Request)

func MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetApiKEY(r)
		if err != nil {
			utils.ResponseWithError(w, 400, err.Error())
			return
		}
		_, err = models.ApiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			utils.ResponseWithError(w, 403, err.Error())
			return
		}
		handler(w, r)
	}
}
