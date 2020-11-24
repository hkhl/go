package article

import (
	"fmt"
	"log"
	"myblog/controllers"
	"myblog/models"
)

type DeleteArticleController struct {
	controllers.BaseController
}

func (this *DeleteArticleController) Get() {
	//根据id删除数据
	id, _ := this.GetInt("id")
	fmt.Println(id)

	_, err := models.DeleteArticleWithId(id)
	if err != nil {
		log.Println(err)
	}
	//重定向
	this.Redirect("/", 302)
}