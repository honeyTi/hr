package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:title,
		Content:content,
		Created:time.Now(),
		Updated:time.Now(),
	}
	_,err := o.Insert(topic)
	return err
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("blog_topic")
	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	}else {
		_,err = qs.All(&topics)
	}
	return topics, err
}