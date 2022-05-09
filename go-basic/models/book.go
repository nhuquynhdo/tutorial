package models

import(
	"example/hello/config"
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB

type Book struct{
	gorm.Model
	BookName string `gorm:""json:"book_name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func RemoveBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

