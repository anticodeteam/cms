package controllers

import (
	"cms/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type SysController struct {
	beego.Controller
}

type Tree struct {
	Id    int `orm:"pk;auto"`
	Title string
	Pid   int
}

func (c *SysController) Tree() {
	dataList, err := models.QueryAllUserInfo()
	if err == nil {
		c.Data["List"] = dataList
	}
	logs.Info("dataList :", dataList)
	c.TplName = "sys_tree.tpl"
}

func (c *SysController) AddTree() {
	title := c.Input().Get("title")
	err := models.InsertAdminDepart(title)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

func (c *SysController) AddTrees() {
	title := c.Input().Get("title")
	err := models.InsertKnowledge(title)
	logs.Info("error:", err)
	c.ServeJSON() //没作用
}

func (c *SysController) GetTree() {
	c.TplName = "page_left.tpl"
}

func (c *SysController) GetPage() {
	dataList, err := models.SearchTree()
	if err == nil {
		c.Data["json"] = dataList
	}
	logs.Info("dataList :", dataList)
	c.ServeJSON()
}

func (c *SysController) Admin() {
	id, _ := c.GetInt("Id") //获取前台传的值
	switch id {
	case 1:
		datalist, _ := models.GetInformationByKonwledge()
		title1 := models.SearchTemplate(id)
		c.Data["BigTitle"] = title1
		c.Data["List"] = datalist
		c.TplName = "user_knowledge.tpl"
	case 2:
		c.TplName = "user_guanzhu.tpl"
	case 3:
		c.TplName = "user_collection.tpl"
	default:
		c.TplName = "blank.tpl"
	}
}

func (c *SysController) EditTree() {
	title := c.Input().Get("title")
	fmt.Println("title:", title)
	pid, _ := strconv.Atoi(c.Input().Get("code"))
	fmt.Println("pid:", pid)
	err := models.EditTrees(title, pid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//删除一级目录
func (c *SysController) DeleteTree() {
	pid, _ := strconv.Atoi(c.Input().Get("code"))
	err := models.DeleteTrees(pid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

