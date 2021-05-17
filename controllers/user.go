package controllers

import (
	"encoding/json"
	"fmt"
	"job/models"
	"job/utils"
	"strconv"
	"time"
)

// UserController operations for User
type UserController struct {
	MainController
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var v models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":mobile")
	t := models.User{}
	err := t.GetUserByMobile(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = t
	}
	c.ServeJSON()
}

func (c *UserController) Login() {
	c.TplName = "user/user_login.html"
}

func (c *UserController) Register() {
	c.TplName = "user/register.html"
}

func (c *UserController) RegisterPost() {
	t := models.User{}
	t.Mobile = c.GetString("mobile")
	t.Password = c.GetString("password")
	t.Nickname = c.GetString("nickName")
	if t.RegisterUser() {
		c.Data["json"] = 1
	} else {
		c.Data["json"] = 0
	}
	c.ServeJSON()
}

func (c *UserController) LoginPost() {
	mobile := c.GetString("userName")
	password := c.GetString("userPass")
	t := &models.User{}
	if mobile == "" || password == "" {
		c.Data["json"] = 0
	}

	if t.LoginUser(mobile, password) {
		t.GetUserByMobile(mobile)
		if utils.Session == nil {
			utils.Session, _ = utils.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		}
		utils.Session.Set("user", t)
		c.Data["json"] = 1
	} else {
		c.Data["json"] = 0
	}
	c.ServeJSON()
}

func (c *UserController) LogOut() {
	// 清楚session
	utils.Session.Delete("user")

	c.Redirect("/user/login", 302)
}

func (c *UserController) Comment() {
	posid, _ := strconv.Atoi(c.GetString("posId"))
	type1, _ := strconv.Atoi(c.GetString("type"))
	content := c.GetString("content")
	user := utils.Session.Get("user").(*models.User)
	comm := models.Comment{
		Type:        type1,
		Content:     content,
		ReleaseTime: time.Now(),
		UserId:      user.Id,
		PositionId:  posid,
	}
	_, result := models.AddComment(&comm)
	if result == nil {
		c.Redirect("/position/"+c.GetString("posId"), 302)
	} else {
		c.Redirect("manager/404.html", 302)
	}
}

func (c *UserController) Info() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "user/user_info.html"
	} else {
		// 获取用户
		user := utils.Session.Get("user").(*models.User)
		if user == nil {
			c.Data["json"] = "fail"
			c.Redirect("/", 302)
		}
		// 获取简历信息
		resume, _ := models.GetResumeByUId(user.Id)
		// 获取收藏职位
		favorPosList, _ := models.ListFavPos(user.Id)
		var applyPOS interface{}
		var applyPosPub interface{}
		if resume != nil {
			// 处理完成记录
			applyPOS, _ = models.ListAppPosHR(resume.Id)
			// 待处理记录
			applyPosPub, _ = models.ListAppPosHRPub(resume.Id)
		}

		// 所有分类记录
		allCate, _ := models.GetAllCategory()

		output := make(map[string]interface{})
		output["user"] = user
		output["resume"] = resume
		output["favorPosList"] = favorPosList
		output["applyPosList"] = applyPOS
		output["prePosList"] = applyPosPub
		output["allCategoryList"] = allCate

		c.Data["json"] = output
		c.ServeJSON()
	}

}
func (c *UserController) ShowResume() {
	if c.Ctx.Request.Method == "GET" {

	} else {
		ret := make(map[string]interface{})
		user := &models.User{}
		resume := &models.Resume{}
		user = utils.Session.Get("user").(*models.User)
		resume, _ = models.GetResumeByUId(user.Id)
		if resume == nil {
			ret["user"] = user
			ret["resume"] = ""
		} else {
			ret["user"] = user
			ret["resume"] = resume
		}
		c.Data["json"] = ret
		c.ServeJSON()
	}

}

func (c *UserController) ResuUpdate() {
	user := utils.Session.Get("user").(*models.User)
	if user == nil {
		c.Redirect("/user/login", 302)
	} else {
		resume := models.Resume{
			Ability:        c.GetString("ability"),
			Internship:     c.GetString("internship"),
			WorkExperience: c.GetString("workExperience"),
			Certificate:    c.GetString("certificate"),
			JobDesire:      c.GetString("jobDesire"),
			UserId:         user.Id,
		}
		if v, _ := models.GetResumeByUId(user.Id); v != nil {
			v.Ability = resume.Ability
			v.Internship = resume.Internship
			v.WorkExperience = resume.WorkExperience
			v.Certificate = resume.WorkExperience
			v.JobDesire = resume.JobDesire
			models.UpdateResumeById(v)
		} else {
			models.AddResume(&resume)
		}
		c.Redirect("/user/info?type=resume", 302)
	}
}

func (c *UserController) UpdateInfo() {
	user := utils.Session.Get("user").(*models.User)
	if user == nil {
		c.Redirect("/user/login", 302)
	}
	nickname := c.GetString("nickname")
	email := c.GetString("email")
	name := c.GetString("name")
	birthYear, _ := strconv.Atoi(c.GetString("birthYear"))
	graYear, _ := strconv.Atoi(c.GetString("graYear"))
	province := c.GetString("province")
	user.City = c.GetString("city")
	eduDegree := c.GetString("eduDegree")
	graduation := c.GetString("graduation")
	major := c.GetString("major")
	dirDesire, _ := strconv.Atoi(c.GetString("dirDesire"))
	user.Nickname = nickname
	user.Email = email
	user.BirthYear = birthYear
	user.Name = name
	user.GraYear = graYear
	user.Province = province
	user.EduDegree = eduDegree
	user.Major = major
	user.Graduation = graduation
	user.DirDesire = dirDesire
	if user.UpdateUser() {
		utils.Session.Delete("user")
		utils.Session.Set("user", user)
		c.Data["json"] = "success"
	}
	c.Data["json"] = "fail"
	c.ServeJSON()
}

func (c *UserController) FavOrNot() {
	user := utils.Session.Get("user").(*models.User)
	id, _ := strconv.Atoi(c.GetString(":id"))
	postion, _ := models.GetPositionById(id)
	if v, _ := models.GetFavorByRePos(user.Id, postion.PosId); v == nil {
		c.Data["json"] = 0
	} else {
		c.Data["json"] = 1
	}
	c.ServeJSON()
}

func (c *UserController) FavPos() {
	user := utils.Session.Get("user").(*models.User)

	id, _ := strconv.Atoi(c.GetString(":id"))
	if user == nil {
		c.Data["json"] = "fail"
	}
	fav := models.Favor{
		UserId:     user.Id,
		PositionId: id,
	}
	_, err := models.AddFavor(&fav)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = "fail"
	} else {
		c.Data["json"] = "success"
	}
	c.ServeJSON()
}

func (c *UserController) DisFavPos() {
	user := utils.Session.Get("user").(*models.User)
	id, _ := strconv.Atoi(c.GetString(":id"))

	err := models.DeleteFavor(user.Id, id)
	if err == nil {
		c.Data["json"] = "success"
	} else {
		c.Data["json"] = "fail"
	}
	c.ServeJSON()
}

func (c *UserController) ApplyOrNot() {
	user := utils.Session.Get("user").(*models.User)
	id, _ := strconv.Atoi(c.GetString(":id"))
	resume, _ := models.GetResumeByUId(user.Id)
	if resume == nil {
		c.Redirect("/user/info?type=resume", 302)
	} else {
		if x, _ := models.GetApplicationByRePos(resume.Id, id); x == nil {
			c.Data["json"] = 1
		} else {
			c.Data["json"] = 0
		}
	}
	c.ServeJSON()
}

func (c *UserController) Apply() {
	user := utils.Session.Get("user").(*models.User)
	id, _ := strconv.Atoi(c.GetString(":id"))
	getPos, _ := models.GetPositionById(id)

	resume, _ := models.GetResumeByUId(user.Id)
	if resume == nil {
		c.Redirect("/user/info?type=resume", 302)
	} else {
		app := models.Application{
			State:      0,
			RecentTime: time.Now(),
			ResumeId:   resume.Id,
			PositionId: id,
			HrId:       getPos.HrIdPub,
		}
		if x, _ := models.GetApplicationByRePos(resume.Id, id); x == nil {
			result, _ := models.AddApplication(&app)
			if result == 0 {
				c.Redirect("manager/404.html", 302)
			} else {
				hr, _ := models.GetHrById(getPos.HrIdPub)
				utils.SendMail(user.Name, user.Mobile, hr.HrEmail, hr.HrName, user.Name, "简历", int(result))
				c.Data["json"] = "1"
			}
		} else {
			c.Data["json"] = "0"
		}
	}
	c.ServeJSON()

}

func (c *UserController) Success() {
	c.TplName = "user/apply_success.html"
}

func (c *UserController) RePass() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "user/RePass.html"
	} else {
		phone := c.GetString("tel")
		user := models.User{}
		user.GetUserByMobile(phone)
		user.Password = utils.Md5Encode(c.GetString("password"))
		if user.UpdateUser() {
			c.Data["json"] = "success"
		} else {
			c.Data["json"] = "failed"
		}
		c.ServeJSON()
	}
}
