package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Collection struct {
	Id          int
	KnowledgeId int
	Uid         int
	Time        string
}

func CollectionInformation(uid interface{}) (dataList []interface{}, err error) {
	var list []Collection
	var list2 []Knowledge //知识库一级
	o := orm.NewOrm()
	qs1 := o.QueryTable(new(Collection))
	qs2 := o.QueryTable(new(Knowledge))
	if _, err = qs1.Filter("uid__exact", uid).All(&list); err == nil { //查询全部
		for _, v := range list {
			if err = qs2.Filter("id__exact", v.KnowledgeId).One(&list2); err == nil {
				for _, values := range list2 { //除了range还有没有其他方法拿到list2里面的值？
					fmt.Println("num:", values)
					dataList = append(dataList, values)
				}
			}
		}
		return dataList, nil
	}
	return nil, err
}

func DeleteCollections(id int, uid interface{}) error {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM collection  WHERE knowledge_id = ? and uid = ?", id, uid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}

func InsertCollection(id int, uid interface{}) error {
	t := time.Now()
	o := orm.NewOrm()
	data := &Collection{}
	data.KnowledgeId = id
	data.Uid = uid.(int)
	data.Time = t.String()
	_, err := o.Insert(data)
	return err
}
