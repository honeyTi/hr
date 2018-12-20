package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	isExit := l.Input().Get("exit")
	if isExit == "true" {
		fmt.Println("执行cookie修改")
		l.Ctx.SetCookie("uname", "",-1, "/")
		l.Ctx.SetCookie("pwd", "",-1, "/")
		l.Redirect("/", 301)
	}
	l.TplName = "login.tpl"
}

func (l *LoginController) Post() {
	uname := l.Input().Get("uname")
	pwd := l.Input().Get("pwd")
	autoLogin := l.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
			maxAge := 0
			if autoLogin {
				maxAge = 1<<31 -1
			}
			l.Ctx.SetCookie("uname", uname, maxAge, "/")
			l.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		l.Redirect("/",301)
		return
	} else {
		l.Redirect("/login",301)
		return
	}

}

func CheckAccount(c *context.Context) bool  {
	ck, err := c.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname:= ck.Value

	pwd1, err := c.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := pwd1.Value

	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd
}