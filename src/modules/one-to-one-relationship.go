package modules

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//OneToOneRelationship : func
func OneToOneRelationship(db *gorm.DB) {

	fmt.Println("============================")
	fmt.Println("1-to-1 Relationship")
	fmt.Println("============================")

	db.CreateTable(&Details{})

	db.Model(&Details{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	//create record
	db.Save(&User{
		NickName:  "ick",
		FirstName: "sdf;slkjdf",
		Balance:   8560,
		Email:     "mail@mail.com",
		Details: Details{
			Link: "testlink",
		},
	})

	u := User{}
	d := Details{}
	db.Last(&u).Related(&d, "details")
	fmt.Println(d)

}
