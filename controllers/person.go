package controllers

import "github.com/astaxie/beego"

type PersonAll struct {
	beego.Controller
}

func (p *PersonAll) Get() {
	p.Data["userName"] = "测试name"
	p.Data["Teb"] = "15010680216"
	p.Data["trueCode"] = true
	p.Data["falseCode"] = false

	type U struct {
		Name string
		Age int
		Sex string
	}

	user := &U{
		Name:"lili",
		Age:12,
		Sex:"女",
	}
	p.Data["user"] = user

	nums := []int{1,2,3,4,5,6,7,8}

	p.Data["nums"] = nums

	p.Data["html"] = "<div>hello beego</div>"

	p.Data["Pipe"] = "<div>hello beego</div>"
}