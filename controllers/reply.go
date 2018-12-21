package controllers

import (
	"beego-map-test/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ReplyControllers struct {
	beego.Controller
}

func (r *ReplyControllers) Add() {
	fmt.Println("ceshi ceshi ")
	tid := r.Input().Get("tid")
	err := models.AddReply(tid, r.Input().Get("nickname"), r.Input().Get("content"))
	if err != nil {
		beego.Error(err)
		return
	}
	r.Redirect("/topic/view/"+tid, 302)
	return
}