package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
	_ "mvc/models"
	_ "mvc/routers"
)

type User struct {
	Uid int
	Pwd string
}

func main() {
	//orm.RegisterDataBase("default","mysql","root:123456@tcp(172.16.1.227:3306)/cmssystem?charset=utf8")
	beego.BConfig.WebConfig.Session.SessionOn = true //开启session
	beego.Run()
}
