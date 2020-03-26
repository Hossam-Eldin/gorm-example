package main

import (
	"gorm-example/src/app"
	"gorm-example/src/database"
	"gorm-example/src/modules"

	"github.com/jinzhu/gorm"
)

func main() {
	db := database.Connect()
	cleanDB(db)
	db.AutoMigrate(&modules.User{}, &modules.Details{}, &modules.Posts{})
	modules.OneToOneRelationship(db)
	app.Routes()

	defer db.Close()
}

func cleanDB(db *gorm.DB) {
	if db.HasTable(&modules.User{}) {
		db.DropTable(&modules.User{})
	}
	if db.HasTable(&modules.Details{}) {
		db.DropTable(&modules.Details{})
	}
	if db.HasTable(&modules.Posts{}) {
		db.DropTable(&modules.Posts{})
	}

}
