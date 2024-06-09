package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open("root:Greeshma@18@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	} else {
		log.Print("Starting server ....")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
