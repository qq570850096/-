package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
	"job/controllers"
	_ "job/routers"
)

func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	// 定时任务
	controllers.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()
	beego.Run()
}
