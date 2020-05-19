package main

import (
	_ "cms/models"
	_ "cms/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid int
	Pwd string
}

func main() {
	//orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/cmssystem?charset=utf8")
	beego.BConfig.WebConfig.Session.SessionOn = true //开启session 的
	beego.Run()
}