package article

import (
	"fmt"
	"myblog/controllers"
	"myblog/models"
)

type UpdateArticleController struct {
	controllers.BaseController
}

func (this *UpdateArticleController) Get() {
	//获取id数据
	id, _ := this.GetInt("id")
	fmt.Println(id)
	//查询数据
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id

	this.TplName = "write_article.html"
}

func (this *UpdateArticleController) Post() {
	//点击提交够的操作
	id, _ := this.GetInt("id")
	fmt.Println(id)
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	//实例化model 修改数据库
	art := models.Article{id, title, tags, short, content, "", 0}
	_, err := models.UpdateArticle(art)

	//返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		response = map[string]interface{}{"code":1, "message":"ok"}
	} else {
		response = map[string]interface{}{"code":0, "message":"error"}
	}
	this.Data["json"] = response
	this.ServeJSON()
}