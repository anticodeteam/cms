package controllers

import (
	"cms/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	str := c.Ctx.Request.RemoteAddr
	fmt.Println(str)
	var user = models.Cms_User{}
	user.Uid, _ = c.Ctx.Input.Session("UserID").(int)

	o := orm.NewOrm()
	err2 := o.Read(&user)
	if err2 == nil {
		fmt.Println("2次登陆 %v ", user)
		if user.Auth == "2" { //Auth为2时，代表登录者为管理员
			c.TplName = "index.tpl"
		} else { //Auth为1时，代表登录者为普通用户
			//c.ctx.Redirect(302, "/loginAction")
			c.TplName = "page_left.tpl"

		}
		return
	}

	c.TplName = "maincontroller/login.tpl"

}

func (c *MainController) Home() {
	c.TplName = "index.tpl"
}

func (c *MainController) Login() {
	var user = models.Cms_User{}
	var Pwd string
	userId, _ := c.GetInt("Uid")
	c.SetSession("UserID", userId)
	Pwd = c.GetString("Pwd")
	user.Uid = userId
	o := orm.NewOrm()
	err2 := o.Read(&user)
	if err2 == nil {
		if user.Pwd == Pwd {
			if user.Auth == "2" { //Auth为2时，代表登录者为管理员
				c.TplName = "index.tpl"
			} else { //Auth为1时，代表登录者为普通用户
				c.TplName = "page_left.tpl"
			}
		} else {
			c.Data["Tip"] = "用户名或密码错误，请重新输入！"
			return
		}
	} else {
		c.Data["Tip"] = "用户名不存在，请重新输入！"
		return
	}
}

func (c *MainController) RegistUser() {
	fmt.Println("-------------------跳转到注册页面------------------")
	c.TplName = "maincontroller/regist.tpl"
}

func (c *MainController) RegistDbInfo() {
	fmt.Println("-------------------开始注册DB信息------------------")
	var user = models.Cms_User{}
	//user struct
	user.Uid, _ = c.GetInt("Uid")
	user.Pwd = c.GetString("Pwd")
	fmt.Println("user.Uid=", user.Uid, " user.Pwd=", user.Pwd)
	//c.SetSession("UserID",userId)
	//跳转到model注册，注册成功跳转回登录页面，注册失败返回当前画面，重新注册
	result := models.RegistDbInformation(user.Uid, user.Pwd)
	if result == "" {
		c.TplName = "maincontroller/login.tpl"
	} else {
		c.Data["result"] = result
		c.TplName = "maincontroller/regist.tpl"
	}
}

func (c *MainController) CancelRegist() {
	c.TplName = "maincontroller/login.tpl"
}

func (c *MainController) LinktoAgent() {
	dateList, err := models.GetAllUserInfo()
	if err == nil {
		c.Data["List"] = dateList
		c.TplName = "maincontroller/AuthAssign.tpl"
	}
}

func (c *MainController) AssignAgent() {
	fmt.Println("---------------------AssignAgentStart-----------------------")
	Uid, _ := c.GetInt("Uid")
	fmt.Println("AssignAgent_Uid=", Uid)
	models.UpdateAgentInfo(Uid)
	c.TplName = "maincontroller/AuthAssign.tpl"
}
