package database

import (
	"log"

	"github.com/jinzhu/gorm"
	//for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Connect : to start db connection
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/gorm_play?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("connection is online")

	return db

}
