package article

import (
	"fmt"
	"myblog/controllers"
	"myblog/models"
	"time"
)

type AddArticleController struct {
	 controllers.BaseController
}

func (this *AddArticleController) Get() {
	this.TplName = "write_article.html"
}

func (this *AddArticleController) Post() {
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model 写入数据库
	art := models.Article{0, title, tags, short, content, "deerwang", time.Now().Unix()}
	_, err := models.AddArticle(art)

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