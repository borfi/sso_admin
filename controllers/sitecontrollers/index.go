package sitecontrollers

import (
	"github.com/astaxie/beego"
)

// SiteController .
type SiteController struct {
	beego.Controller
}

// Index .
func (c *SiteController) Index() {
	c.Data["Website"] = "Index page"
	c.Data["Email"] = "index@gmail.com"
	c.TplName = "site/index.tpl"
}

// Test .
func (c *SiteController) Test() {
	mystruct := map[string]interface{}{"name": "test", "age": 33}
	c.Data["json"] = &mystruct
	c.ServeJSON()
	c.TplName = "site/index.tpl"
}
