package controllers

import (
	"fmt"
	"job/models"
	"job/utils"
	"strconv"
	"strings"
	"time"
)

type HrController struct {
	MainController
}

func (c *HrController) Login() {
	c.TplName = "hr/login.html"
}

func (c *HrController) MailBox() {
	c.TplName = "hr/mailbox.html"
}
func (c *HrController) LoginPost() {
	out := make(map[string]interface{})
	out["state"] = "0"
	username := c.GetString("username")
	password := c.GetString("password")
	if v, flag := models.LoginHr(username, password); flag == true {
		out["state"] = "1"
		if utils.Session == nil {
			utils.Session, _ = utils.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		}
		utils.Session.Set("hr", v)
	}
	c.Data["json"] = out
	c.ServeJSON()
}

func (c *HrController) GetPosInfo() {

	c.TplName = "hr/popform.html"
}

func (c *HrController) Index() {
	if utils.Session == nil {
		c.Redirect("/hr/login", 302)
	}
	c.TplName = "hr/index.html"
}
func (c *HrController) Indv3() {
	c.TplName = "hr/index_v3.html"
}
func (c *HrController) HrInfo() {
	c.TplName = "hr/nestable_list.html"
}
func (c *HrController) PostInfo() {
	out := make(map[string]interface{})
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		v := utils.Session.Get("hr").(*models.Hr)
		id := v.Id
		// 收件箱
		applist, _ := models.ListAppByHr(id)
		//创建的职位
		poslist, _ := models.ListPositionByHr(id)
		out["hr"] = v
		out["applyPosList"] = applist
		out["positions"] = poslist
		c.Data["json"] = out
	}
	c.ServeJSON()
}

func (c *HrController) GetInfo() {
	out := make(map[string]interface{})
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		v := utils.Session.Get("hr").(*models.Hr)
		out["hr"] = v
		c.Data["json"] = out
	}
	c.ServeJSON()
}

func (c *HrController) InfoUpdate() {
	hr := utils.Session.Get("hr").(*models.Hr)
	hr.HrPassword = c.GetString("password")
	hr.HrName = c.GetString("name")
	hr.HrEmail = c.GetString("email")
	hr.DepartmentId, _ = strconv.Atoi(c.GetString("departmentId"))
	if err := models.UpdateHrById(hr); err == nil {
		c.Redirect("hr/index.html", 302)
	} else {
		c.Data["json"] = err
		c.ServeJSON()
	}

}

func (c *HrController) Logout() {
	if err := utils.Session.Delete("hr"); err == nil {
		c.Redirect("hr/login", 302)
	}
}

func (c *HrController) ApplyUser() {
	if utils.Session.Get("hr") == nil {
		c.Data["json"] = "用户认证已过期或未登录，请登录hr账号"
	} else {
		id, _ := strconv.Atoi(c.GetString(":id"))
		if v, err := models.GetApplicationById(id); err == nil {
			v.State = 1
			models.UpdateApplicationById(v)
			c.Data["json"] = "success"
		} else {
			fmt.Println(err)
			c.Data["json"] = "fail"
		}
	}

	c.ServeJSON()
}

func (c *HrController) DeUser() {
	if utils.Session.Get("hr") == nil {
		c.Data["json"] = "用户认证已过期或未登录，请登录hr账号"
	} else {
		id, _ := strconv.Atoi(c.GetString(":id"))
		if v, err := models.GetApplicationById(id); err == nil {
			v.State = 2
			models.UpdateApplicationById(v)
			c.Data["json"] = "success"
		} else {
			fmt.Println(err)
			c.Data["json"] = "fail"
		}
	}
	c.ServeJSON()
}

func (c *HrController) PosList() {
	page, _ := strconv.Atoi(c.GetString("page"))
	limit, _ := strconv.Atoi(c.GetString("limit"))
	if utils.Session.Get("hr") == nil {
		c.Data["json"] = "用户认证已过期或未登录，请登录hr账号"
	} else {
		ret1 := make(map[string]interface{})
		v := utils.Session.Get("hr").(*models.Hr)
		pos, _ := models.HrPos(v.Id)
		ret1["count"] = len(pos)
		if len(pos) >= page*limit {
			pos = pos[page*limit-limit : page*limit]
		} else {
			pos = pos[page*limit-limit:]
		}
		flag := false
		title := ""
		data := make([]interface{}, 0)
		if c.GetString("key[title]") != "" {
			flag = true
			title = strings.ToLower(c.GetString("key[title]"))
		}

		for _, x := range pos {
			cate, _ := models.GetCategoryById(x.CategoryId)
			depa, _ := models.GetDepartmentById(x.DepartmentId)
			ret := make(map[string]interface{})
			ret["id"] = x.PosId
			ret["title"] = x.Title
			ret["quantity"] = x.Quantity
			ret["workCity"] = x.WorkCity
			ret["SalaryDown"] = x.SalaryDown
			ret["SalaryUp"] = x.SalaryUp
			ret["requirement"] = x.Requirement
			ret["department"] = depa.DepartmentName
			ret["category"] = cate.CategoryName
			if flag == true {
				tmp := strings.ToLower(x.Title)
				if strings.Contains(tmp, title) {
					data = append(data, ret)
				}
			} else {
				data = append(data, ret)
			}

		}

		ret1["data"] = data

		ret1["msg"] = ""
		ret1["code"] = 0
		c.Data["json"] = ret1
	}
	c.ServeJSON()
}

func (c *HrController) PosDel() {
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		id, _ := strconv.Atoi(c.GetString("id"))
		if err := models.DeletePosition(id); err == nil {
			c.Data["json"] = "success"
		} else {
			fmt.Println(err)
			c.Data["json"] = "fail"
		}
	}

	c.ServeJSON()
}

func (c *HrController) HrDepart() {
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		hr := utils.Session.Get("hr").(*models.Hr)
		depart, _ := models.GetDepartmentById(hr.DepartmentId)
		comp, _ := models.GetCompanyById(depart.CompanyId)
		ret := make(map[string]interface{})
		list, _ := models.GetDepartByCompany(comp.Id)
		ret["data"] = list
		ret["comp"] = comp.CompanyName
		ret["user"] = hr
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

func (c *HrController) PostEdit() {
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		id, _ := strconv.Atoi(c.GetString("id"))
		pos, _ := models.GetPositionById(id)
		ret := make(map[string]interface{})
		ret["posInfo"] = pos
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

func (c *HrController) GetCate() {
	ret := make(map[string]interface{})
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		id, _ := strconv.Atoi(c.GetString("id"))
		if id != 0 {
			pos, _ := models.GetPositionById(id)
			cate, _ := models.GetCategoryById(pos.CategoryId)
			ret["Tcategory"] = cate
		}
		all_cate, _ := models.GetAllCategory()
		ret["cate_all"] = all_cate
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

func (c *HrController) PostUpdate() {
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		id, _ := strconv.Atoi(c.GetString("id"))
		pos, _ := models.GetPositionById(id)
		pos.WorkCity = c.GetString("city-picker")
		pos.Title = c.GetString("title")
		pos.Requirement = c.GetString("requirement")
		pos.Quantity, _ = strconv.Atoi(c.GetString("quantity"))
		pos.SalaryDown, _ = strconv.Atoi(c.GetString("salaryDown"))
		pos.SalaryUp, _ = strconv.Atoi(c.GetString("salaryUp"))
		pos.CategoryId, _ = strconv.Atoi(c.GetString("cate"))
		if err := models.UpdatePositionById(pos); err == nil {
			c.Data["json"] = "success"
		} else {
			c.Data["json"] = "fail"
		}
	}
	c.ServeJSON()
}

func (c *HrController) PostAdd() {
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		pos := models.Position{}
		pos.WorkCity = c.GetString("city-picker")
		pos.Title = c.GetString("title")
		pos.Requirement = c.GetString("requirement")
		pos.Quantity, _ = strconv.Atoi(c.GetString("quantity"))
		pos.SalaryDown, _ = strconv.Atoi(c.GetString("salaryDown"))
		pos.SalaryUp, _ = strconv.Atoi(c.GetString("salaryUp"))
		hr := utils.Session.Get("hr").(*models.Hr)
		depart, _ := models.GetDepartmentById(hr.DepartmentId)
		comp, _ := models.GetCompanyById(depart.CompanyId)
		if x, err := models.GetDepartByCompany(comp.Id); err == nil {
			pos.DepartmentId = x[0].Id
		}
		pos.CategoryId, _ = strconv.Atoi(c.GetString("cate"))
		pos.ReleaseDate = time.Now()
		pos.StatePub = 1
		pos.HrIdPub = hr.Id
		if _, err := models.AddPosition(&pos); err == nil {
			c.Data["json"] = "success"
		} else {
			c.Data["json"] = "fail"
		}
	}
	c.ServeJSON()
}

func (c *HrController) GetApply() {
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		var user *models.User

		data := make([]interface{}, 0)
		hrId := utils.Session.Get("hr").(*models.Hr).Id
		if bo, err := models.ListAppByHrId(hrId); err != nil {
			fmt.Println(err)
		} else {
			for _, v := range bo {
				resu, _ := models.GetResumeById(v.ResumeId)
				user, _ = models.GetUser(resu.UserId)
				ret := make(map[string]interface{})
				ret["username"] = user.Name
				ret["email"] = user.Email
				ret["degree"] = user.EduDegree
				ret["Title"] = v.Title
				ret["WorkCity"] = v.WorkCity
				ret["mobile"] = user.Mobile
				ret["id"] = user.Id
				data = append(data, ret)
			}
			c.Data["json"] = data
			c.ServeJSON()
		}
	}
}

func (c *HrController) Check() {
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		id, _ := strconv.Atoi(c.GetString(":id"))
		user, err := models.GetUser(id)
		if err != nil {
			c.Data["json"] = 0
		}
		hrId := utils.Session.Get("hr").(*models.Hr).Id
		if bo, err := models.ListAppByHrId(hrId); err != nil {
			c.Data["json"] = 0
			fmt.Println(err)
		} else {
			for _, v := range bo {
				app, _ := models.GetApplicationById(v.AppId)
				resu, _ := models.GetResumeById(v.ResumeId)
				user_app, _ := models.GetUser(resu.UserId)
				if user.Id == user_app.Id {
					app.State = 3
				}
				models.UpdateApplicationById(app)
				c.Data["json"] = 1
				c.Redirect("/hr/success", 302)
			}
		}
	}
	c.ServeJSON()
}

func (c *HrController) UnCheck() {
	if utils.Session.Get("hr") == nil {
		c.Redirect("/hr/login", 302)
	} else {
		id, _ := strconv.Atoi(c.GetString(":id"))
		user, err := models.GetUser(id)
		if err != nil {
			c.Data["json"] = 0
		}
		hrId := utils.Session.Get("hr").(*models.Hr).Id
		if bo, err := models.ListAppByHrId(hrId); err != nil {
			fmt.Println(err)
			c.Data["json"] = 0
		} else {
			for _, v := range bo {
				app, _ := models.GetApplicationById(v.AppId)
				resu, _ := models.GetResumeById(v.ResumeId)
				user_app, _ := models.GetUser(resu.UserId)
				if user.Id == user_app.Id {
					app.State = 4
				}
				models.UpdateApplicationById(app)
				c.Data["json"] = 1
				c.Redirect("/hr/success", 302)
			}
		}
	}
	c.ServeJSON()
}

func (c *HrController) Success() {
	c.TplName = "hr/hr_success.html"
}
