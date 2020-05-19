package controllers

import (
	"cms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "maincontroller/login.tpl"
}

func (c *MainController) Home() {
	c.TplName = "index.tpl"
}

func (c *MainController) Login() {
	var user = models.User{}
	var Pwd string
	userId, _ := c.GetInt("Uid")
	user.Uid = userId
	Pwd = c.GetString("Pwd")
	c.SetSession("UserID", userId)
	o := orm.NewOrm()
	err2 := o.Read(&user)
	if err2 == nil {
		if user.Pwd == Pwd {
			c.TplName = "index.tpl"
		} else {
			c.Data["Tip"] = "用户名或密码错误，请重新输入！"
			return
		}
	} else {
		c.Data["Tip"] = "用户名不存在，请重新输入！"
		return
	}

}
