// @APIVersion 1.0.0
// @Title myuser API
// @Description A user system for multi-platforms.
// @Contact hsowan.me@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License MIT
// @LicenseUrl https://github.com/hsowan-me/myuser/blob/master/LICENSE
package routers

import (
	"myuser/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
