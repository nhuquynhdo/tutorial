package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"example/hello/models"
	"example/hello/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	fmt.Println("bookId: ", bookId)
	ID, err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("GetBookById::error when parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()

	res, _ := json.Marshal(b)
	w.Write(res)
}

func RemoveBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("RemoveBook::error when parsing")
	}
	book := models.RemoveBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("UpdateBook::error when parsing")
	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.BookName != "" {
		bookDetails.BookName = updateBook.BookName
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}