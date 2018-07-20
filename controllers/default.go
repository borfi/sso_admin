package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Test Test"
	c.Data["Email"] = "Test@gmail.com"
	c.TplName = "index.tpl"
}
