package controllers

import (
	"beego-map-test/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
	topics, err  := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
	c.TplName = "index.tpl"
}
