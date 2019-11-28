package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "myuser/routers"

	"github.com/astaxie/beego"
)

func init()  {
	fmt.Println(beego.AppConfig.String("sqlconn"))
	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"), 30)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
