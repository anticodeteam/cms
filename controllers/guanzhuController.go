package controllers

import (
	"cms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type GuanzhuController struct {
	beego.Controller
}

type Jsons struct {
	flag string "Json"
}

//关注一览界面
func (c *GuanzhuController) GetUserGuanzhu() {
	uid := c.GetSession("UserID")
	datalist, err := models.GuanzhuInformation(uid)
	if err == nil {
		c.Data["json"] = datalist
	}
	logs.Info("dataList :", datalist)
	c.ServeJSON()
}

//点击添加关注
func (c *GuanzhuController) AddGuanzhuInformation() {
	userId := c.GetSession("UserID")
	id, _ := c.GetInt("Id")
	pid, _ := c.GetInt("Pid")
	isinsert := models.AddGuanzhuInformation(userId, id, pid)
	if isinsert != 0 {
		c.Data["json"] = map[string]interface{}{"flag": true}
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//取消关注
func (c *GuanzhuController) DeleteGuanzhu() {
	id, _ := c.GetInt("Id")
	uid := c.GetSession("UserID")
	err := models.DeleteGuanzhuInfo(id, uid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}
