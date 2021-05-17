package controllers

import (
	"github.com/astaxie/beego"
	"job/models"
	"job/utils"
	"strconv"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	c.TplName = "user/search_list.html"
}
func (c *SearchController) PostSearch() {
	keyword := c.GetString("keyword") + "%"
	orderby := c.GetString("orderBy")
	if orderby == "" {
		orderby = "salaryUp"
	}
	page, _ := strconv.Atoi(c.GetString("page"))
	limit := 6
	posInfo, _ := models.ListSearchPos(keyword, orderby)
	var pagetotal int
	if cap(posInfo) >= limit*page {
		pagetotal = len(posInfo) / 6
		if len(posInfo)%6 != 0 {
			pagetotal++
		}
		posInfo = posInfo[page*limit-limit : page*limit]
	} else {
		pagetotal = 1
	}
	out := make(map[string]interface{})
	out["user"] = utils.Session.Get("user")
	out["title"] = "第" + c.GetString("page") + "页"
	out["keyword"] = keyword
	out["orderBy"] = orderby
	out["posInfo"] = posInfo
	out["pagetotal"] = pagetotal
	c.Data["json"] = out
	c.ServeJSON()

}
