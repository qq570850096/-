package controllers

import (
	"github.com/astaxie/beego"
	"job/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	if utils.Session == nil {
		c.Redirect("/user/login", 302)
	}

	c.TplName = "user/index.html"
}

func (c *MainController) SearchList() {
	c.TplName = "user/search_list.html"
}
