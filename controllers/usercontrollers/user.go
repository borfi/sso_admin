package usercontrollers

import "github.com/astaxie/beego"

// UserController .
type UserController struct {
	beego.Controller
}

// List .
func (c *UserController) List() {
	c.Data["Website"] = "User page"
	c.Data["Email"] = "user@gmail.com"
	c.TplName = "user/list.tpl"
}
