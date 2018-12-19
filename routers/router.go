package routers

import (
	"beego-map-test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/index", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{})
}
