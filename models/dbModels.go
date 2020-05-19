package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//设置DB信息                    别名 必须有一个default               数据库类型            用户：密码@连接方式（一般是tcp）/数据库名称？编码
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(172.16.1.66:3306)/ly_db?charset=utf8")
	//映射Model数据          建表
	orm.RegisterModel(new(Cms_User), new(Cms_Tree), new(Cms_Knowledge), new(Cms_Guanzhu), new(Cms_Collection), new(Cms_Article))
	//生成表              别名        是否强制更新      是否可见（创建过程）
	orm.RunSyncdb("default", false, false)
}
