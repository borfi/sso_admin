package main

import (
	_ "sso_admin/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Info(beego.BConfig.AppName, "1.0")
	beego.Run()
}
