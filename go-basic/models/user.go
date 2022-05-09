package models

import(
	"example/hello/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct{
	gorm.Model
	UserID int `gorm:""json:"user_id"`
	UserName string `json:"user_name"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User{
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User{
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(ID int64) (*User, *gorm.DB) {
	var u User
	db := db.Where("ID=?", ID).Find(&u)
	return &u, db
}

func RemoveUser(ID int64) User{
	var u User
	db.Where("ID=?", ID).Delete(u)
	return u
}