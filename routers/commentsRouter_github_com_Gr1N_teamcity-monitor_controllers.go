package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/Gr1N/teamcity-monitor/controllers:APIController"] = append(beego.GlobalControllerRouter["github.com/Gr1N/teamcity-monitor/controllers:APIController"],
		beego.ControllerComments{
			"Builds",
			`/builds`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Gr1N/teamcity-monitor/controllers:APIController"] = append(beego.GlobalControllerRouter["github.com/Gr1N/teamcity-monitor/controllers:APIController"],
		beego.ControllerComments{
			"BuildsStatus",
			`/builds/status`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Gr1N/teamcity-monitor/controllers:APIController"] = append(beego.GlobalControllerRouter["github.com/Gr1N/teamcity-monitor/controllers:APIController"],
		beego.ControllerComments{
			"BuildsStatistics",
			`/builds/statistics`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Gr1N/teamcity-monitor/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/Gr1N/teamcity-monitor/controllers:IndexController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

}
