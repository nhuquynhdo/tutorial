package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	db * gorm.DB
)

func ConnectDB(){
	d, err := gorm.Open("mysql", "root:Abc@123654@(127.0.0.1:3308)/HelloSql?parseTime=true&charset=utf8&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
	fmt.Println("Successfully Connected to MySQL database")
	
}

func GetDB() *gorm.DB{
	return db
}