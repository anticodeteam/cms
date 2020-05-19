package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Knowledge struct {
	Id         int
	Title      string
	Pid        int
	Gid        int
	UpdateTime string
	Creater    string
}

type Knowledge2 struct {
	Id         int
	Title      string
	Pid        int
	Gid        int
	Isguanzhu  int
	UpdateTime string
	Creater    string
}

type Article struct {
	Id          int
	KnowledgeId int
	Title       string
	Detail      string
	Creater     string
	Time        string
}

func GetInformationByKonwledge() (dataList []interface{}, err error) {
	var list []Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Knowledge))
	//查询
	//查询数据
	if _, err = qs.Filter("pid__exact", 0).Filter("gid__exact", 0).All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func InsertDirectory(title string, id int) error {
	o := orm.NewOrm()
	category := &Knowledge{}
	category.Title = title
	category.Pid = 0
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func SearchKnowledge(pid int) (dataList2 []interface{}, err error) {
	var list []Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Knowledge))
	if pid == 0 {
		return nil, err
	}
	//查询
	//查询数据
	if _, err = qs.Filter("Pid__exact", pid).All(&list); err == nil {
		for _, v := range list {
			dataList2 = append(dataList2, v)
		}
		return dataList2, nil
	}
	return nil, err
}

func GetDirectorys() (dataList []interface{}, err error) {
	var list []Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Knowledge))
	//查询
	//查询数据
	if _, err = qs.Filter("pid__exact", 0).Filter("gid__exact", 0).All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func GetMoreDirectorys(title string) (dataList []interface{}, err error) {
	var list []Knowledge
	var know = new(Knowledge)
	o := orm.NewOrm()
	qs := o.QueryTable(new(Knowledge))
	o.Raw("select * from knowledge where  title = ?", title).QueryRow(&know)
	//查询
	//查询数据
	//if _, err = qs.Filter("pid__exact", 0).Filter("gid__exact", 0).All(&list); err == nil {
	if _, err = qs.Filter("pid__exact", know.Id).All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

//测试自建结构体插入
func Knowledges(userid interface{}) (dataList []interface{}, err error) {

	var list []Knowledge
	var data = new(Knowledge2)
	o := orm.NewOrm()
	qs := o.QueryTable(new(Knowledge))
	//if _, err = qs.Filter("pid__exact", 0).All(&list); err == nil {            只查询一级目录
	if _, err = qs.All(&list); err == nil { //查询全部
		for _, v := range list {

			if userid == nil {
				data = &Knowledge2{v.Id, v.Title, v.Pid, v.Gid, 0, v.UpdateTime, v.Creater} //自定义输入内容，前台根据isguanzhu来判定是否显示取消关注，还差一个查询关注表的内容
			} else {

				data = &Knowledge2{v.Id, v.Title, v.Pid, v.Gid, IsGuanzhu(userid.(int), v.Id), v.UpdateTime, v.Creater} //自定义输入内容，前台根据isguanzhu来判定是否显示取消关注，还差一个查询关注表的内容
			}

			dataList = append(dataList, *data)
		}
		fmt.Println("datalist12321321:", dataList)
		return dataList, nil
	}
	return nil, err
}

func InsertKnowledge(title string) error {
	o := orm.NewOrm()
	category := &Knowledge{}
	category.Title = title
	category.Pid = 0
	category.Creater = "顾康"
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func JumpToKnowledgePage() (dataList []interface{}, err error) {
	var list []Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Knowledge))
	if _, err = qs.Filter("pid__exact", 0).Filter("gid__exact", 0).All(&list); err == nil { //只查询一级目录
		//if _, err = qs.All(&list); err == nil {										//查询全部
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, err
	}
	return nil, err
}

func EditKnowledge(title string, pid int) error {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE knowledge SET title = ? WHERE id = ?", title, pid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}

func DeleteKnowledge(pid int) error {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM knowledge  WHERE id = ?", pid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}

func InsertArticle(text, title string, id int) error {
	t := time.Now() //设置当前时间
	o := orm.NewOrm()
	category := &Article{}
	data := &Knowledge{}
	data.Title = title
	data.Pid = 0
	data.Creater = "顾康"
	data.Gid = id
	data.UpdateTime = t.String()
	dataId, _ := o.Insert(data)
	category.KnowledgeId = int(dataId)
	category.Title = title
	category.Detail = text
	category.Time = t.String() //将时间转换成string类型进行保存
	category.Creater = "顾康"
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func GetArticles(id int) (dataList []interface{}, err error) {
	var list []Article
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	if _, err = qs.Filter("knowledge_id", id).All(&list); err == nil {
		//if _, err = qs.All(&list); err == nil { //查询全部
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func GetArticleLists() (dataList []interface{}, err error) {
	var list []Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Knowledge))
	if _, err = qs.Filter("pid__exact", 0).Filter("gid__gt", 0).All(&list); err == nil { //只查询一级目录
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func JumpAllKnowPage(id int) (dataList []interface{}, err error) {
	var list []Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Knowledge))
	if _, err = qs.Filter("pid__exact", 0).Filter("gid__exact", id).All(&list); err == nil { //只查询一级目录
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}
