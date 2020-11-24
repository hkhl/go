package main

import (
	"github.com/astaxie/beego"
	_ "myblog/routers"
	"myblog/utils"
)

func main() {
	utils.InitMysql()
	//beego.BConfig.WebConfig.Session.SessionOn = true //打开session
	beego.Run()
}

