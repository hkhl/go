package controllers

import (
	"fmt"
	"myblog/models"
)

type AlbumController struct {
	BaseController
}
func (this *AlbumController) Get() {
	albums,err := models.FindAllAlbums()
	if err !=nil{
		fmt.Println(err)
	}
	this.Data["Album"] = albums
	this.TplName="album.html"
}