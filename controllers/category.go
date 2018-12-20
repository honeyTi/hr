package controllers

import (
	"beego-map-test/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	c.Data["Category"] = true
	c.Data["IsLogin"] = CheckAccount(c.Ctx)

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category",301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
	}
	var err error
	c.Data["Categories"], err = models.GetAllCategory()
	if err != nil {
		fmt.Println(err)
	}
	c.TplName = "category.tpl"
}