package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Cms_Guanzhu struct {
	Id          int
	KnowledgeId int
	Uid         int
	Pid         int
	Time        string
}

func GuanzhuInformation(uid interface{}) (dataList []interface{}, err error) {
	var list []Cms_Guanzhu
	var list2 []Cms_Knowledge //知识库一级
	var list3 []Cms_Knowledge //知识库二级
	o := orm.NewOrm()
	qs1 := o.QueryTable(new(Cms_Guanzhu))
	qs2 := o.QueryTable(new(Cms_Knowledge))
	//if _, err = qs.Filter("pid__exact", 0).All(&list); err == nil {            只查询一级目录
	if _, err = qs1.Filter("uid__exact", uid).All(&list); err == nil { //查询全部
		for _, v := range list {
			if err = qs2.Filter("id__exact", v.KnowledgeId).One(&list2); err == nil {
				for _, values := range list2 { //除了range还有没有其他方法拿到list2里面的值？
					dataList = append(dataList, values)
					if _, err = qs2.Filter("Pid__exact", values.Id).All(&list3); err == nil {
						for _, val := range list3 {
							dataList = append(dataList, val)
						}
					}
				}
			}
		}
		return dataList, nil
	}
	return nil, err
}

func AddGuanzhuInformation(userId interface{}, id, pid int) (isinsert int64) {
	var data Cms_Guanzhu
	o := orm.NewOrm()
	data.KnowledgeId = id
	data.Pid = pid
	data.Uid = userId.(int)
	isinsert, err := o.Insert(&data)
	if err == nil {
		return
	}
	return
}

func DeleteGuanzhuInfo(id int, uid interface{}) error {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM guanzhu  WHERE knowledge_id = ? and uid = ?", id, uid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}
