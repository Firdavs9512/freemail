package databases

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db variable for initialization base
func Connect() *gorm.DB {
	dsn := "firdavs:nobody@tcp(127.0.0.1:3000)/freemail?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error for connection databases!")
	}
	return db
}

var db *gorm.DB

func BaseInit() {
	db = Connect()
}
