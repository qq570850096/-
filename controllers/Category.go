package controllers

import (
	"github.com/astaxie/beego"
	"job/models"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

// 职位分类页分页输出 （职位分类，职位列表）
func (c *CategoryController) CateList() {
	id, _ := strconv.Atoi(c.GetString(":id"))
	page, _ := strconv.Atoi(c.GetString(":page"))
	cate, _ := models.GetCategoryById(id)
	if cate == nil {
		c.TplName = "/manager/404.html"
	} else {
		posInfo, _ := models.ListCatePos(id)
		if cap(posInfo) >= page*12 {
			posInfo = posInfo[page*12-12 : page*12]
		}
		out := make(map[string]interface{})
		out["title"] = "第" + c.GetString(":page") + "页"
		out["category"] = cate
		out["posInfo"] = posInfo
		c.Data["json"] = out
		c.ServeJSON()
	}
}
