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
		c.Redirect("/login", 302)
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")
	tid := c.Input().Get("tid")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}

func (c *TopicController) View() {
	c.TplName = "topic_view.tpl"
	topic, err := models.GetTopicById(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = c.Ctx.Input.Param("0")
}

func (t *TopicController) Modify() {
	t.TplName = "topic_modify.tpl"
	tid := t.Input().Get("tid")
	topic, err := models.GetTopicById(tid)
	if err != nil {
		beego.Error(err)
		t.Redirect("/", 302)
		return
	}
	t.Data["Topic"] = topic
	t.Data["Tid"] = tid
}

func (t *TopicController) Delete() {
	if !CheckAccount(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopicMap(t.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	t.Redirect("/", 302)
}