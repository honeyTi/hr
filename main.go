package main

import (
	"beego-map-test/models"
	_ "beego-map-test/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.InitDB()

	orm.RegisterDriver("mysql", orm.DRMySQL)

	database_url := beego.AppConfig.String("databaseUser") + ":" + beego.AppConfig.String("databasePass") +
		"@tcp(" + beego.AppConfig.String("databaseUrls") + ":" +beego.AppConfig.String("databasePort") +
		")/" + beego.AppConfig.String("databaseDB") + "?charset=utf8"

	fmt.Println(database_url)

	orm.RegisterDataBase("default", beego.AppConfig.String("databaseType"), database_url,50)
}

func main() {
	// 开启orm调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	beego.Run()
}