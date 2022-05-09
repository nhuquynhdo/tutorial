package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"example/hello/routes"
)

func main(){
	r := mux.NewRouter()
	routes.BookRoutes(r)
	routes.UserRoutes(r)
	log.Fatal(http.ListenAndServe(":8001", r))
}
