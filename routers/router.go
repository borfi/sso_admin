package routers

import (
	"fmt"
	"sso_admin/controllers"

	beecontext "github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	var auth beego.FilterFunc = func(ctx *beecontext.Context) {
		if ua := ctx.Input.UserAgent(); ua != "11" {
			fmt.Println("################ auth ###########################")
			return
		}
		fmt.Println("###########EEEEEEEEEEEE auth EEEEEEEEEEEEEEEEEEEEEEEEEEEE")
		return
	}

	ns := beego.NewNamespace("/v1",
		beego.NSCond(func(ctx *beecontext.Context) bool {
			if ua := ctx.Input.UserAgent(); ua != "11" {
				fmt.Println("###########################################")
				return true
			}
			fmt.Println("###########EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")
			return false
		}),
		beego.NSBefore(auth),
		beego.NSRouter("/t", &controllers.MainController{}),
		beego.NSNamespace("/test",
			beego.NSRouter("/test", &controllers.MainController{}),
		),
	)

	beego.AddNamespace(ns)
}
