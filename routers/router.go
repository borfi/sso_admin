package routers

import (
	"sso_admin/controllers/sitecontrollers"
	"sso_admin/controllers/usercontrollers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &sitecontrollers.SiteController{}, "get:Index")

	beego.Router("/site/list", &sitecontrollers.SiteController{}, "get:Index")
	beego.Router("/site/test", &sitecontrollers.SiteController{}, "get:Test")

	beego.Router("/user/list", &usercontrollers.UserController{}, "get:List")
}
