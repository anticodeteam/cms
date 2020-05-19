package controllers

import (
	"cms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CollectionController struct {
	beego.Controller
}

func (c *CollectionController) GetUserCollection() {
	uid := c.GetSession("UserID")
	datalist, err := models.CollectionInformation(uid)
	if err == nil {
		c.Data["json"] = datalist
	}
	logs.Info("dataList :", datalist)
	c.ServeJSON()
}

func (c *CollectionController) DeleteCollection() {
	id, _ := c.GetInt("Id")
	uid := c.GetSession("UserID")
	err := models.DeleteCollections(id, uid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//添加收藏
func (c *CollectionController) AddCollection() {
	uid := c.GetSession("UserID")
	id, _ := c.GetInt("articleid")
	err := models.InsertCollection(id, uid)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}
