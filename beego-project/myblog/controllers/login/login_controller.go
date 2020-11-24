package login

import (
	"fmt"
	_ "fmt"
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
)


type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, "password:", password)
	id := models.QueryUserWithParam(username, utils.MD5(password))
	if id > 0 {
		//设置session后, 带上cookie
		this.SetSession("loginuser", username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}