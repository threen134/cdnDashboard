package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// func (c *MainController) Get() {
// 	c.Layout = "layout.tpl"
// 	c.LayoutSections = make(map[string]string)
// 	c.LayoutSections["PageContent"] = "instances.tpl"
// 	c.LayoutSections["SideBare"] = "sidebar.tpl"
// 	c.TplName = "instances.tpl"
// }

// // @router /index [get]
// func (c *MainController) Index() {
// 	c.Layout = "layout.tpl"
// 	c.LayoutSections = make(map[string]string)
// 	c.LayoutSections["PageContent"] = "instances.tpl"
// 	c.LayoutSections["SideBare"] = "sidebar.tpl"
// 	c.TplName = "instances.tpl"
// }

// @router  /instances [get]
func (c *MainController) Instances() {
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["PageContent"] = "instances.tpl"
	c.LayoutSections["SideBar"] = "sidebar.tpl"
	c.LayoutSections["CSS"] = "css.tpl"
	c.LayoutSections["JS"] = "js.tpl"
	c.TplName = "instances.tpl"
}
