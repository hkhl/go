package article

import (
	"myblog/controllers"
	"myblog/models"
	"myblog/utils"
	"strconv"
)

type ShowArticleController struct {
	controllers.BaseController
}

// "/article/:id"
func (this *ShowArticleController) Get() {
	//获取id数据
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//查询数据
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	this.TplName = "show_article.html"
}