package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Tree struct {
	Id    int `orm:"pk;auto"`
	Title string
	Pid   int
}

func QueryAllUserInfo() (dataList []interface{}, err error) {
	var list []Tree
	o := orm.NewOrm()
	qs := o.QueryTable(new(Tree))
	//查询
	//查询数据
	if _, err = qs.Filter("Pid__exact", 0).All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func InsertAdminDepart(title string) error {
	o := orm.NewOrm()
	category := &Tree{}
	category.Title = title
	category.Pid = 0
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
	//res, err := o.Raw("Insert into tree (title,pid) values (?,?) ","测试",0).Exec()
	//if err == nil {
	//	num, _ := res.RowsAffected()
	//	fmt.Println("mysql row affected nums: ", num)
	//	return err
	//}
	//return err
}

func InsertAdminDepart3(title string, pid int) error {
	o := orm.NewOrm()
	category := &Knowledge{}
	category.Title = title
	category.Pid = pid
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

//func QueryAllUserInfo2(pid int) (dataList2 []interface{}, err error) {
//	var list []Tree
//	o := orm.NewOrm()
//	qs := o.QueryTable(new(Tree))
//	fmt.Println("models的pid:",pid)
//	if pid == 0 {
//		//fmt.Println("没值")
//		return nil, err
//	}
//	//查询
//	//查询数据
//	if _, err = qs.Filter("Pid__exact", pid).All(&list); err == nil {
//		for _, v := range list {
//			dataList2 = append(dataList2, v)
//		}
//		//fmt.Println("models:",dataList2)
//		return dataList2, nil
//	}
//	return nil, err
//}

func SearchTree() (dataList []interface{}, err error) {
	var list []Tree
	o := orm.NewOrm()
	qs := o.QueryTable(new(Tree))
	//查询
	//查询数据
	_, err = qs.All(&list)
	if err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func SearchTemplate(id int) string {
	o := orm.NewOrm()
	t1 := new(Tree) //一级目录集合
	t2 := new(Tree) //二级目录集合
	o.Raw("select * from tree where  id = ?", id).QueryRow(&t2)
	if t2.Pid == 0 {
		return t2.Title //返回一级目录的标题及ID
	}
	o.Raw("select * from tree where  id = ?", t2.Pid).QueryRow(&t1)
	return t1.Title //T1标题是一级目录标题  T2标题是二级目录标题
}

func EditTrees(title string, pid int) error {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE tree SET title = ? WHERE id = ?", title, pid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}

func DeleteTrees(pid int) error {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM tree  WHERE id = ?", pid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}
