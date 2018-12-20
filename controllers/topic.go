package controllers

import (
	"beego-map-test/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["Topic"] = true
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
	topics, err  := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
	c.TplName = "topic.tpl"
}

func (c *TopicController) Add() {
	c.TplName = "topic_add.tpl"
}

func (c *TopicController) Post() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("./login", 302)
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")

	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}