package main

import (
	"log"
	"net/http"

	"github.com/Hayoun01/book_store_api/pkg/config"
	"github.com/Hayoun01/book_store_api/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./pkg/config/.env")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	v1Router := r.PathPrefix("/v1").Subrouter()
	routes.RegisterBookRoutes(v1Router)
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    ":" + config.Port,
	}
	log.Println("Server running on", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
