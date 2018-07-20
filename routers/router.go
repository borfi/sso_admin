package routers

import (
	"sso_admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/t", &controllers.MainController{}),
		beego.NSNamespace("/test",
			beego.NSRouter("/test", &controllers.MainController{}),
		),
	)

	beego.AddNamespace(ns)
}
