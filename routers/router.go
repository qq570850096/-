// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"job/controllers"

	"github.com/astaxie/beego"
)

func init() {

	// sms 请求控制
	beego.Router("/sms/sendCode", &controllers.SMSController{}, "post:SendCode")
	beego.Router("/sms/verifyCode", &controllers.SMSController{}, "post:VerifyCode")
	beego.Router("/sms/genexcel", &controllers.SMSController{}, "get:GenExcel")
	// user 请求控制
	beego.Router("/user/register", &controllers.UserController{}, "get:Register")
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/user/registerPost", &controllers.UserController{}, "post:RegisterPost")
	beego.Router("/user/loginPost", &controllers.UserController{}, "post:LoginPost")
	beego.Router("/user/logout", &controllers.UserController{}, "get:LogOut")
	beego.Router("/user/info", &controllers.UserController{}, "*:Info")
	beego.Router("/user/resume", &controllers.UserController{}, "*:ShowResume")
	beego.Router("/user/comment", &controllers.UserController{}, "post:Comment")
	beego.Router("/user/resume/update", &controllers.UserController{}, "post:ResuUpdate")
	beego.Router("/user/info/update", &controllers.UserController{}, "post:UpdateInfo")
	beego.Router("/user/favorOrNot/:id", &controllers.UserController{}, "get:FavOrNot")
	beego.Router("/user/favor/:id", &controllers.UserController{}, "get:FavPos")
	beego.Router("/user/disfavor/:id", &controllers.UserController{}, "get:DisFavPos")
	beego.Router("/user/apply/:id", &controllers.UserController{}, "*:Apply")
	beego.Router("/user/applyOrNot/:id", &controllers.UserController{}, "*:ApplyOrNot")
	beego.Router("/user/RePass", &controllers.UserController{}, "*:RePass")
	beego.Router("/user/success", &controllers.UserController{}, "*:Success")
	// 分页控制
	beego.Router("/page/:id", &controllers.PageController{})
	beego.Router("/page/genrec", &controllers.PageController{}, "get:GenMGO")
	beego.Router("/page/uprec", &controllers.PageController{},"get:UpdateMGO")
	beego.Router("/category/:id/:page", &controllers.CategoryController{}, "post:CateList")

	beego.Router("/search", &controllers.SearchController{}, "post:PostSearch")
	beego.Router("/searchlist", &controllers.SearchController{}, "get:Get")

	beego.Router("/", &controllers.MainController{})

	// 文件控制
	beego.Router("/upload", &controllers.FileController{}, "post:Upload")
	beego.Router("/read/resume", &controllers.FileController{}, "get:Read")

	// 职位控制
	beego.Router("/position/:id/:page", &controllers.PosController{}, "post:PostPos")
	beego.Router("/position/:id", &controllers.PosController{}, "get:GetPos")

	// 简历管理
	beego.Router("/hr/apply/:id", &controllers.HrController{}, "get:ApplyUser")
	beego.Router("/hr/deny/:id", &controllers.HrController{}, "get:DeUser")
	beego.Router("/hr/login", &controllers.HrController{}, "get:Login")
	beego.Router("/hr/index_v3", &controllers.HrController{}, "get:Indv3")
	beego.Router("/hr/loginPost", &controllers.HrController{}, "get:LoginPost")
	beego.Router("/hr/index", &controllers.HrController{}, "get:Index")
	beego.Router("/hr/info", &controllers.HrController{}, "post:PostInfo")
	beego.Router("/hr/info", &controllers.HrController{}, "get:GetInfo")
	beego.Router("/hr/info/update", &controllers.HrController{}, "post:InfoUpdate")
	beego.Router("/logout", &controllers.HrController{}, "post:Logout")
	beego.Router("/hr/mailbox", &controllers.HrController{}, "get:MailBox")
	beego.Router("/hr/positionShow", &controllers.HrController{}, "get:PosList")
	beego.Router("/hr/nestable_list", &controllers.HrController{}, "get:HrInfo")
	beego.Router("/hr/depart", &controllers.HrController{}, "get:HrDepart")
	beego.Router("/hr/getPosInfo", &controllers.HrController{}, "get:GetPosInfo")
	beego.Router("/hr/posEdit", &controllers.HrController{}, "get:PostEdit")
	beego.Router("/hr/cate_pos", &controllers.HrController{}, "get:GetCate")
	beego.Router("/hr/pos_delete", &controllers.HrController{}, "post:PosDel")
	beego.Router("/hr/pos_update", &controllers.HrController{}, "post:PostUpdate")
	beego.Router("/hr/posAdd", &controllers.HrController{}, "post:PostAdd")
	beego.Router("/hr/app_list", &controllers.HrController{}, "get:GetApply")
	beego.Router("/hr/check/:id", &controllers.HrController{}, "get:Check")
	beego.Router("/hr/uncheck/:id", &controllers.HrController{}, "get:UnCheck")
	beego.Router("/hr/success", &controllers.HrController{}, "get:Success")
}
