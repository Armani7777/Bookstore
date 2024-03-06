package config

import (
	"github.com/Armani7777/Bookstore/app"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func Init() *gorm.DB {
	dsn := "root:root@/productsdb?parseTime=true"
	conn, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalln(err)
	}
	db = conn
	db.AutoMigrate(&app.Product{}, &app.Cart{}, &app.Fav{})

	return db
}
func GetDB() *gorm.DB {
	return db
}
