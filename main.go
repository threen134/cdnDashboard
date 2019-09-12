package main

import (
	"github.com/astaxie/beego"
	_ "github.com/threen134/cdnDashboard/routers"
)

func main() {
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateRight = ">>>"
	beego.BConfig.WebConfig.AutoRender = true

	beego.Run()
}
