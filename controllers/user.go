package controllers

import (
	"fmt"
	"myuser/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err == nil {
		fmt.Println(user)
		if uid, e := models.AddUser(user); e == nil {
			u.Data["json"] = map[string]int64{"uid": uid}
		} else {
			u.Data["json"] = e.Error()
		}
	} else {
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}

// @Title Get the user
// @Description get user by uid
// @Param	uid		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid, err := u.GetInt64(":uid")
	if err == nil {
		user, e := models.GetUserById(uid)
		if e != nil {
			u.Data["json"] = e.Error()
		} else {
			u.Data["json"] = user
		}
	} else {
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}

// @Title Update the user
// @Description update the user
// @Param	uid		path 	int64	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router / [put]
func (u *UserController) Put() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if models.UpdateUser(&user) {
		u.Data["json"] = "Update success!"
	} else {
		u.Data["json"] = "Update fail!"
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	if uid, err := u.GetInt64(":uid"); err == nil {
		if models.DeleteUser(uid) {
			u.Data["json"] = "Delete success!"
		} else {
			u.Data["json"] = "Delete fail!"
		}
	} else {
		u.Data["json"] = "Bad request!"
	}
	u.ServeJSON()
}

// @Title Auth
// @Description Auth
// @Param	principle		query 	string	true		"The principle for auth"
// @Param	password		query 	string	true		"The password for auth"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /auth [post]
func (u *UserController) Auth() {
	principle := u.GetString("principle")
	password := u.GetString("password")
	if authType, err := u.GetInt("authType"); err == nil {
		if models.Auth(principle, password, authType) {
			u.Data["json"] = "login success!"
		} else {
			u.Data["json"] = "login fail!"
		}
	} else {
		u.Data["json"] = "Error auth type!"
	}
	u.ServeJSON()
}

