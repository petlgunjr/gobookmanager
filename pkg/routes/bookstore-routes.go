package routes

import (
	"github.com/gorilla/mux"
	"github.com/petlgunjr/gobookmanager/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandlerFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandlerFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
