package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

// @router /index [get]
func (c *MainController) Index() {
	c.TplName = "index.html"
}

// @router  /instances [get]
func (c *MainController) Instances() {
	c.TplName = "instances.html"
}
