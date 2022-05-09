package routes

import (
	"example/hello/controllers"
	"github.com/gorilla/mux"
	"fmt"
)

var BookRoutes = func(router *mux.Router) {
	fmt.Println("-- BookRoutes")
	router.HandleFunc("/books/add", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.RemoveBook).Methods("DELETE")
	
}