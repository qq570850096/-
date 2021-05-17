package controllers

import (
	"github.com/astaxie/beego"
	"job/models"
	"job/utils"
	"strconv"
)

type PosController struct {
	beego.Controller
}

func (c *PosController) GetPos() {
	c.TplName = "user/position_detail.html"
}

// 职位细节页 评论分页输出 （职位，部门，公司，分类，评论列表）
func (c *PosController) PostPos() {
	id, _ := strconv.Atoi(c.GetString(":id"))
	page, _ := strconv.Atoi(c.GetString(":page"))
	position, err := models.GetPositionById(id)
	user := utils.Session.Get("user").(*models.User)
	hr := utils.Session.Get("hr")
	if err != nil {
		c.Redirect("/manager/404.html", 302)
	}
	position.Hits++
	models.UpdatePositionById(position)
	output := make(map[string]interface{})
	//所属部门信息
	department, _ := models.GetDepartmentById(position.DepartmentId)
	company, _ := models.GetCompanyById(department.CompanyId)
	cate, _ := models.GetCategoryById(position.CategoryId)
	comlist := models.ListCommentPage(position.PosId, page, 6)
	output["position"] = position
	output["department"] = department
	output["company"] = company
	output["category"] = cate
	output["comList"] = comlist
	if user != nil {
		output["user"] = user
	} else if hr != nil {
		output["hr"] = hr
	}
	c.Data["json"] = output
	c.ServeJSON()
}
