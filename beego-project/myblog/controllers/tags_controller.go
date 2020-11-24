package controllers

import (
	"fmt"
	"myblog/models"
)

type TagsController struct {
	BaseController
}


func (this *TagsController) Get() {
	tagsString := models.QueryArticleWithParam("tags")
	fmt.Println("tags:= ", tagsString)
	tagsList := models.HandleTagsListData(tagsString)
	this.Data["Tags"] = tagsList
	this.TplName = "tags.html"
}