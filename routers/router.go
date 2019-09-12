package routers

import (
	"github.com/astaxie/beego"
	"github.com/threen134/cdnDashboard/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("index.html", &controllers.MainController{}, "get:Index")
	beego.Router("instances.html", &controllers.MainController{}, "get:Instances")
}
