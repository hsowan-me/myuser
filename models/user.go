package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// register model
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int
	Username string
	Password string
	Email string
	Phone string
	Avatar string
}

func AddUser(u User) string {
	return ""
}

func GetUser(uid string) (u *User, err error) {

	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return nil
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	return nil, errors.New("User Not Exist")
}

func Auth(username, password string) bool {

	return false
}

func DeleteUser(uid string) {

}
