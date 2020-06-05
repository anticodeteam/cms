package controllers

import (
	"cms/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"strconv"
)

type KnowController struct {
	beego.Controller
}

//func (c *KnowController) AddDirectory(){
//	title := c.Input().Get("title")
//	id, _ := strconv.Atoi(c.Input().Get("code"))
//	err := models.InsertDirectory(title,id)
//	logs.Info("dataList :", err)
//	if err != nil {
//		c.ServeJSON()
//	}else{
//		c.ServeJSON()
//	}
//}

//func (c *KnowController) SearchKnowledge()  {
//	pid, _ := strconv.Atoi(c.Input().Get("pid"))
//	dataList2, err2 := models.SearchKnowledge(pid)
//	if err2 == nil {
//		c.Data["json"] =dataList2
//	}
//	logs.Info("dataList2 :", dataList2)
//	c.ServeJSON()
//}

//func (c *KnowController) Getdirectory() {
//	pid, _ := strconv.Atoi(c.Input().Get("pid"))
//	var datas []interface{}
//	switch pid {
//	case 1: datalist , _ := models.GetDirectorys()
//		datas = datalist
//	case 2:
//	case 3:
//	}
//	c.Data["json"] =datas
//	logs.Info("dataList2 :", datas)
//	c.ServeJSON()
//}

//后台页面点击一级知识库目录出现二级目录
func (c *KnowController) GetMoreDirectory() {
	title := c.GetString("title")
	datalist, err := models.GetMoreDirectorys(title)
	if err == nil {
		c.Data["json"] = datalist
	}
	logs.Info("dataList :", datalist)
	c.ServeJSON()
}

//拼写前端页面知识库画面
func (c *KnowController) GetKnowledge() {
	userId := c.GetSession("UserID")
	datalist, err := models.Knowledges(userId)
	if err == nil {
		c.Data["json"] = datalist
	}
	logs.Info("dataList :", datalist)
	c.ServeJSON()
}

//跳转知识总页面
func (c *KnowController) JumpPage() {
	fid, _ := c.GetInt("id")
	//c.Data["gid"] = fid
	//c.TplName = "knowledge.tpl"
	datalist, err := models.JumpAllKnowPage(fid)
	if err == nil {
		c.Data["Fid"] = fid
		c.Data["List"] = datalist
		c.TplName = "knowledge.tpl"
	} else {
		c.Data["Fid"] = fid
		c.Data["List"] = datalist
		c.TplName = "knowledge.tpl"
	}
}

//跳转知识详细页面
func (c *KnowController) JumpArticlePage() {
	id, _ := c.GetInt("id")
	c.Data["articleId"] = id
	c.TplName = "knowledge_article.tpl"
}

//跳转知识库管理页面
func (c *KnowController) JumpKnowledgePages() {
	datalist, err := models.JumpToKnowledgePage()
	if err == nil {
		c.Data["List"] = datalist
	}
	c.TplName = "back_knowledgePage.tpl"
}

//添加知识库二级目录
func (c *KnowController) AddMoreTrees() {
	title := c.Input().Get("title")
	pid, _ := strconv.Atoi(c.Input().Get("code"))
	err := models.InsertAdminDepart3(title, pid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//编辑一级目录
func (c *KnowController) EditKnowledge() {
	title := c.Input().Get("title")
	pid, _ := strconv.Atoi(c.Input().Get("code"))
	err := models.EditKnowledge(title, pid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//删除一级目录
func (c *KnowController) DeleteKnowledge() {
	pid, _ := strconv.Atoi(c.Input().Get("code"))
	err := models.DeleteKnowledge(pid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//跳转添加知识文章页面
func (c *KnowController) JumpAddArticle() {
	gid, _ := c.GetInt("id")
	c.Data["gid"] = gid
	c.TplName = "addArticle.tpl"
}

//添加知识文章
func (c *KnowController) AddArticle() {
	texts := c.GetString("text")
	title := c.GetString("title")
	id, _ := c.GetInt("id")
	Uid := c.GetSession("UserID")
	err := models.InsertArticle(texts, title, id, Uid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//显示知识文章
func (c *KnowController) GetArticle() {
	articleidId, _ := c.GetInt("articleid")
	datalist, err := models.GetArticles(articleidId)
	if err == nil {
		c.Data["json"] = datalist
	}
	//logs.Info("dataList :", datalist)
	c.ServeJSON()
}

func (c *KnowController) Jumptest() {
	c.TplName = "test.tpl"
}

//初始化文章列表
func (c *KnowController) GetArticleList() {
	datalist, err := models.GetArticleLists()
	if err == nil {
		c.Data["json"] = datalist
	}
	//logs.Info("dataList :", datalist)
	c.ServeJSON()
}

//用户页面上传知识等待审批
func (c *KnowController) UserSaveKonwledge() {
	knowledgeName := c.GetString("Name")
	uid := c.GetSession("UserID")
	fmt.Println(knowledgeName, uid)
	err := models.UserSaveKonwledgeAction(knowledgeName, uid)
	logs.Info("err:", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//管理员修改知识点审批状态
func (c *KnowController) ChangeKnowledgeStatus() {
	id, _ := c.GetInt("id")
	status, _ := c.GetInt("status")
	models.ChangeKnowledgeStatusAction(id, status)
	c.TplName = "knowledge.tpl"
}

//跳转上传页面
func (c *KnowController) Jumpopload() {
	titleid, _ := c.GetInt("id")
	c.Data["titleid"] = titleid
	c.TplName = "maincontroller/upload.tpl"
}

//文件上传
func (c *KnowController) UploadFile() {

	file, head, err := c.GetFile("file")
	if err != nil {
		c.Ctx.WriteString("获取文件失败")
		return
	}
	defer file.Close()
	fmt.Println(file)
	filename := head.Filename
	err = c.SaveToFile("file", "SDPATH/"+filename)
	if err != nil {
		c.Ctx.WriteString("上传失败1")
	} else {
		c.Ctx.WriteString("上传成功")
	}

}
func (c *KnowController) Addfilename() {
	//filename := c.UploadFile("fiflename")
	filename := c.GetString("filename")
	titleid, _ := c.GetInt("titleid")
	fmt.Println("filename=", filename)
	err := models.SaveFileName(filename, titleid)
	//logs.Info("dataList",err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}

}

//获取文件名
func (c *KnowController) Getfilename() {
	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(`D:\gopath\src\cms\SDPATH`)
	if err != nil {
		panic(err)
	}
	var filename []string
	// 获取文件，并输出它们的名字
	for _, file := range files {
		//fmt.Println(file.Name())
		name := file.Name()
		filename = append(filename, name)
		fmt.Println(filename)
	}
	c.Data["json"] = filename
	c.ServeJSON()
}

//下载文件
func (c *KnowController) Download() {
	filename := c.GetString("filename")
	fmt.Println("filename=", filename)
	src := "./SDPATH/" + filename
	fmt.Println("src=", src)
	c.Ctx.Output.Download(src)
	//c.Ctx.Output.Download("SDPATH/HKZ-go.xlsx")
}

//评论
func (c *KnowController) JumpComments() {
	id, _ := c.GetInt("id")
	c.TplName = "maincontroller/comments.tpl"
	fmt.Println("knowledgeid=", id)
	c.Data["konwledgeId"] = id
}

//获取评论
func (c *KnowController) GetComments() {
	//knowledgeid, _ := c.GetInt("id")
	//读取表中现有的评论
	//var dataList interface{}
	//dataList,err := models.Getcomments()
	//if err != nil{
	//	beego.Error(err)
	//}
	//id := c.JumpComments()
	id, _ := c.GetInt("id")
	//fmt.Println("koneledgeID+",id)
	datalist, err := models.GetCommentList(id)
	fmt.Println("datalist=", datalist)
	//c.TplName="maincontroller/comments.tpl"
	if err == nil {
		c.Data["json"] = datalist
	}
	c.ServeJSON()
	//c.TplName="maincontroller/comments.tpl"
}

//保存评论
func (c *KnowController) Savecomments() {
	id, _ := c.GetInt("id")
	name := c.GetString("name")
	comment := c.GetString("comment")
	fmt.Println(id, name, comment)
	models.SaveComments(id, name, comment)
	c.Data["konwledgeId"] = id
	c.TplName = "maincontroller/comments.tpl"
}

//删除评论
func (c *KnowController) DeleteComment() {
	name := c.GetString("name")
	fmt.Println("name====", name)
	models.DeleteComment(name)
	c.TplName = "maincontroller/comments.tpl"
}

//改变知识状态
func (c *KnowController) ChangeThisStatus() {
	Id, _ := c.GetInt("Id")
	status, _ := c.GetInt("status")
	models.ChangeThisStatusAction(Id, status)
	c.ServeJSON()
}

//
func (c *KnowController) AddLevel2Menu() {
	ID, _ := c.GetInt("ID")
	Title := c.GetString("Title")
	Uid := c.GetSession("UserID")
	models.AddLevel2MenuAction(ID, Title, Uid)
	c.ServeJSON()
}

//
func (c *KnowController) DeleteKnow() {
	ID, _ := c.GetInt("ID")
	models.DeleteKnowAction(ID)
	c.ServeJSON()
}
