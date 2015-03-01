package routers

import (
	"github.com/astaxie/beego"

	"github.com/Gr1N/teamcity-monitor/controllers"
)

func init() {
	beego.Include(&controllers.IndexController{})

	ns :=
		beego.NewNamespace("/api",
			beego.NSNamespace("/v1",
				beego.NSInclude(
					&controllers.APIController{},
				),
			),
		)

	beego.AddNamespace(ns)
}
