package routers

import (
	"beego-map-test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
    beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/topic", &controllers.TopicController{})
    beego.Router("/reply", &controllers.ReplyControllers{})
}
