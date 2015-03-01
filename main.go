package main

import (
	"github.com/astaxie/beego"

	_ "github.com/Gr1N/teamcity-monitor/routers"
)

func main() {
	beego.Run()
}
