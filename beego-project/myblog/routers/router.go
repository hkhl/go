package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
	"myblog/controllers/article"
	"myblog/controllers/login"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/register", &login.RegisterController{})
    beego.Router("/login", &login.LoginController{})
    beego.Router("/exit", &controllers.ExitController{})

    //add article
    beego.Router("/article/add", &article.AddArticleController{})
    //show article
    beego.Router("/article/:id", &article.ShowArticleController{})
    //update article
    beego.Router("/article/update", &article.UpdateArticleController{})
    //deleta article
	beego.Router("/article/delete", &article.DeleteArticleController{})

    //tags
    beego.Router("/tags", &controllers.TagsController{})

	//相册
	beego.Router("/album", &controllers.AlbumController{})
	//文件上传
	beego.Router("/upload", &controllers.UploadController{})

    //aboutme
    beego.Router("/aboutme", &controllers.AboutMeController{})
}
