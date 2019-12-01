package models

import (
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"crypto/md5"
)

var o orm.Ormer

const (
	auth_username = 0
	auth_phone = 1
	auth_email = 2
)

func init() {

	//orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"), 30)
	orm.Debug, _ = beego.AppConfig.Bool("ormdebug")

	orm.RegisterModel(new(User))
	o = orm.NewOrm()
}

type User struct {
	Id       int64
	Username string
	Password string
	Email string
	Phone string
	Avatar string
	Salt string
}

type Student struct {
	Id int64
}

type Company struct {
	Id int64
}

func AddUser(u User) (uid int64, err error) {
	return o.Insert(&u)
}

func GetUserById(uid int64) (u *User, err error) {
	user := User{Id: uid}
	e := o.Read(&user)
	return &user, e
}

func UpdateUser(u *User) bool {
	if num, err := o.Update(&u); err == nil {
		return num == 1
	} else {
		return false
	}
}

// @Param	t	auth type	0username	1phone	2email
func Auth(principle string, password string, authType int) bool {
	var user User
	if authType == auth_username {
		user = User{Username: principle}
	} else if authType == auth_phone {
		user = User{Phone: principle}
	} else if authType == auth_email {
		user = User{Email: principle}
	} else {
		return false
	}
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询
	if err := o.Read(&user, "Username"); err == nil {
		h := md5.New()
		h.Write([]byte(password + user.Salt))
		passwd := hex.EncodeToString(h.Sum(nil))
		if passwd == user.Password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func DeleteUser(uid int64) bool {
	if _, err := o.Delete(&User{Id: uid}); err == nil {
		return true
	} else {
		return false
	}
}
