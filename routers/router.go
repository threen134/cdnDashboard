package routers

import (
	"github.com/astaxie/beego"
	"github.com/threen134/cdnDashboard/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Instances")
	beego.Router("index.html", &controllers.MainController{}, "get:Instances")
	beego.Router("instances.html", &controllers.MainController{}, "get:Instances")
}
