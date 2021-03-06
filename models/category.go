package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)
// 添加分类
func AddCategory(name string) error {
	fmt.Println(123213)
	o := orm.NewOrm()
	cate := &Category{Title:name}
	qs := o.QueryTable("blog_category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	fmt.Println("插入数据")
	_, err1 := o.Insert(cate)
	if err1 != nil {
		return err1
	}
	return nil
}
// 获取所有分类
func GetAllCategory() ([] *Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("blog_category")
	_, err := qs.All(&cates)
	return cates, err
}
// 删除分类
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id:cid}
	_, err1 := o.Delete(cate)
	return err1
}