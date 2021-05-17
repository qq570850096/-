package controllers

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego/toolbox"
	"github.com/zuijinbuzai/excelclaim/excel"
	"job/common"
	"job/models"
	"job/utils"
	"log"
	"strconv"
)

type SMSController struct {
	MainController
}

func (c *SMSController) SendCode() {
	str := c.GetString("mobile")
	v, err := utils.SendCode(str)
	if err == nil {
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *SMSController) VerifyCode() {
	phone := c.GetString("mobile")
	code := c.GetString("code")
	v, err := utils.VerifyCode(phone, code)
	if err == true {
		c.Data["json"] = 1
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *SMSController) GenExcel() {
	GenExcel()

	c.Data["json"] = 1

	c.ServeJSON()
}
func InitTask() {
	tk := toolbox.NewTask("generateWarning", "0 15 10 15 * ?", GenExcel)
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("generateWarning", tk)
}
func GenExcel() (err error) {
	userlist, err := models.ListAppByState(3)
	if err != nil {
		log.Fatal(err)
	}
	f := excelize.NewFile()
	sheet := excel.NewSheet(f, "已入职学生", 10, 28)
	sheet.SetAllColsWidth(7, 14, 11, 15, 12, 13, 10, 9, 12, 13)
	excelStyle := excel.NewExcelStyle(11, 0, false)
	sheet.ApplyRows(excelStyle, 6)

	excelStyle2 := excel.NewExcelStyle(11, 0, true)

	sheet.WriteRow("入职学生信息明细").SetRowHeight(34)
	sheet.WriteRow("序号", "姓名", "性别", "手机号", "毕业学校", "专业", "毕业年份", "就业公司", "职位", "薪资")
	sheet.ApplyRowsRange(excelStyle2, 1, 2)
	var gender string
	for i, v := range userlist {
		if v.Gender == 0 {
			gender = "男"
		} else {
			gender = "女"
		}
		str := strconv.Itoa(v.SalaryDown/1000) + "k - " + strconv.Itoa(v.SalaryUp/1000) + "k"
		sheet.WriteRow(strconv.Itoa(i+1), v.Name, gender, v.Mobile, v.Graduation, v.Major, strconv.Itoa(v.GraYear), v.CompanyName, v.Title, str)

	}

	if err := f.SaveAs("./tmp/recruit.xlsx"); err != nil {
		fmt.Println(err)
	}
	if userlist == nil {
		log.Println("暂无入职者，无需发送")
		return
	} else {
		err = utils.SendMailToSchool("网站管理员", common.School_Email, "大连大学", "入职情况统计")
		if err != nil {
			log.Fatal(err)
		}
	}
	return err
}
