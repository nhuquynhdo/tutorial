package routes

import (
	"example/hello/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/register1", controllers.GetBooks).Methods("GET")
}