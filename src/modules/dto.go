package modules

import "github.com/jinzhu/gorm"

var users []User = []User{
	User{FirstName: "Foo", NickName: "Dor", Balance: 200, Email: "emil@mail.com"},
	User{FirstName: "Boo", NickName: "Bor", Balance: 800, Email: "emil@gmail.com"},
	User{FirstName: "Zoo", NickName: "Zor", Balance: 500, Email: "emil@hotmail.com"},
}

//User : dto
type User struct {
	gorm.Model
	Balance int
	//set column type
	FirstName string `sql:"type:VARCHAR(255)"`

	//set deafult value for column
	NickName string `sql:"DEAFULT :'zoro'"`
	//Not Null  & Unique Filed
	Email string `sql:"not null;unique"`

	//ignord example
	Status bool `sql:"-"`

	//Relationship
	//has it won id in other table
	Details Details
	//the user id and the posts id exists in it's won tableas map
	Posts []Posts `gorm:"many2many:auth_user"`
}

//Details : dto
type Details struct {
	ID     uint
	Link   string
	UserID uint
}

//Posts : dto
type Posts struct {
	gorm.Model
	Tilte   string
	Context string
	OwenrID uint
	Author  []User `gorm:"many2many:auth_user"`
}

//NotFound : check if recored not exists in the db
func (u *User) NotFound() bool {
	return u.Model.ID == 0
}
