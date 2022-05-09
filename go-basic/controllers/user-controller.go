package controllers

import(
	"encoding/json"
	// "fmt"
	// "github.com/gorilla/mux"
	"net/http"
	// "strconv"
	"example/hello/models"
	"example/hello/utils"
)

// var NewBook models.Book
func Register(w http.ResponseWriter, r *http.Request){
	user := &models.User{}
	utils.ParseBody(r, user)
	b := user.CreateUser()

	res, _ := json.Marshal(b)
	w.Write(res)
}

