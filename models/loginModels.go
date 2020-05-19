package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Cms_User struct {
	Uid int `orm:"pk;"`
	Pwd string
	Auth string
}

func RegistDbInformation(Uid int,Pwd string)(result string){
	if Uid != 0 {
		o := orm.NewOrm()
		user := Cms_User{Uid:Uid,Pwd:Pwd}
		if created,id,err := o.ReadOrCreate(&user,"uid");err == nil {
			if created {
				fmt.Println("New Insert an Object, uid ",id)
			}else {
				fmt.Println("Get an Object, uid ",id)
				result = "已注册"
			}
		}
	}else{
		result = "用户名必须为工号"
	}
	return result
}

func GetAllUserInfo()(dataList []interface{},err error){
	var list []Cms_User
	var data = new(Cms_User)
	o  := orm.NewOrm()
	qs := o.QueryTable(new(Cms_User))
	if _,err = qs.All(&list);err == nil {
		for _,v := range list {
			data = &Cms_User{v.Uid,v.Pwd,v.Auth}
			dataList = append(dataList, *data)
		}
		return dataList,nil
	}
	return nil, err
}

func UpdateAgentInfo(uid int) (){
	fmt.Println("---------------------UpdateAgentInfo------------------")
	o := orm.NewOrm()
	user := Cms_User{Uid:uid}
	if o.Read(&user) == nil {
		user.Auth = "1"
		if num, err := o.Update(&user);err == nil {
			fmt.Println(num)
		}
	}
}