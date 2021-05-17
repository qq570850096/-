package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"job/models"
	"job/utils"
	"os"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Upload() {
	user := utils.Session.Get("user").(*models.User)
	ret := make(map[string]string)
	file, _, err := c.Ctx.Request.FormFile("file")
	if err != nil {
		fmt.Printf("Failed to get data, err:%s\n", err.Error())
		ret["filefalg"] = "fall"
	}
	defer file.Close()
	location := "./tmp/" + user.Name + "-" + user.Mobile + ".pdf"
	newfile, err := os.Create(location)
	if err != nil {
		fmt.Println("Failed to create file", err)
		ret["filefalg"] = "fall"
	}
	defer newfile.Close()
	_, err = io.Copy(newfile, file)
	if err != nil {
		fmt.Println("Failed to save data into file", err)
		ret["filefalg"] = "fall"
	} else {
		ret["filefalg"] = "success"
	}

	c.Data["json"] = ret
	c.ServeJSON()
}

func (c *FileController) Read() {
	user := utils.Session.Get("user").(*models.User)
	file, err := ioutil.ReadFile("./tmp/" + user.Name + "-" + user.Mobile + ".pdf")
	if err != nil {
		fmt.Println("Failed to save data into file", err)
	} else {
		c.Ctx.ResponseWriter.Write(file)
	}
}
