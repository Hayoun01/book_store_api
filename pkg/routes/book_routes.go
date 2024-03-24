package routes

import (
	"github.com/Hayoun01/book_store_api/pkg/auth"
	"github.com/Hayoun01/book_store_api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(route *mux.Router) {
	route.HandleFunc("/ready", controllers.HandleGetReady).Methods("GET")
	route.HandleFunc("/create_book", auth.MiddlewareAuth(controllers.HandleCreateBook)).Methods("POST")
	route.HandleFunc("/books", controllers.HandleGetAllBooks).Methods("GET")
	route.HandleFunc("/create_author", auth.MiddlewareAuth(controllers.HandleCreateAuthor)).Methods("POST")
	route.HandleFunc("/authors", controllers.HandleGetAllAuthors).Methods("GET")
	route.HandleFunc("/book/{book_id}", controllers.HandleGetBookByID).Methods("GET")
	route.HandleFunc("/book/{book_id}", auth.MiddlewareAuth(controllers.HandleDeleteBookByID)).Methods("DELETE")
	route.HandleFunc("/users", controllers.HandleCreateUser).Methods("POST")
}
