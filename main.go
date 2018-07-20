package main

import (
	_ "sso_admin/routers"
	"strings"

	"github.com/astaxie/beego"
	beegoCtx "github.com/astaxie/beego/context"
)

func main() {
	//beego.Info(beego.BConfig.AppName, "1.0")

	var filterDeal = func(ctx *beegoCtx.Context) {
		loginDeal(ctx)
		authDeal(ctx)
	}
	beego.InsertFilter("/*", beego.BeforeExec, filterDeal)

	beego.Run()
}

//判断是否已经登陆
func loginDeal(ctx *beegoCtx.Context) {
	url := ctx.Input.URL()
	isAjax := ctx.Input.IsAjax()
	account, ok := ctx.Input.Session("account").(string)
	data := map[string]interface{}{"succ": 0, "msg": "报歉，请重新登陆"}
	if !strings.HasPrefix(url, "/site/") {
		if !ok || "" == account {
			if isAjax {
				ctx.Output.JSON(data, false, false)
			} else {
				ctx.Redirect(302, "/site/login")
			}
		}
	} else if strings.HasPrefix(url, "/site/login") {
		if ok && "" != account {
			if isAjax {
				data["msg"] = "不能重复登陆"
				ctx.Output.JSON(data, false, false)
			} else {
				ctx.Redirect(302, "/")
			}
		}
	}
}

//判断是否有权限
func authDeal(ctx *beegoCtx.Context) {
	url := ctx.Input.URL()
	isAjax := ctx.Input.IsAjax()
	data := map[string]interface{}{"succ": 0, "msg": "报歉，您没有此操作权限！"}
	if !strings.HasPrefix(url, "/site/") && "/" != url {
		if !isAjax {
			ctx.Output.JSON(data, false, false)
		}
	}
}
