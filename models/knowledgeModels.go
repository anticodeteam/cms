package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Cms_Knowledge struct {
	Id         int
	Title      string
	Pid        int
	Gid        int
	UpdateTime string
	Creater    string
	Status     int
}

type Cms_Knowledge2 struct {
	Id         int
	Title      string
	Pid        int
	Gid        int
	Isguanzhu  int
	UpdateTime string
	Creater    string
	Status     int
}

type Cms_Article struct {
	Id          int
	KnowledgeId int
	Title       string
	Detail      string
	Creater     string
	Time        string
}

func GetInformationByKonwledge() (dataList []interface{}, err error) {
	var list []Cms_Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))
	//查询
	//查询数据 pid==0，代表一级目录
	if _, err = qs.Filter("pid__exact", 0).Filter("gid__exact", 0).All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		fmt.Println("This datalist--->", dataList)
		return dataList, nil
	}
	return nil, err
}

func InsertDirectory(title string, id int) error {
	o := orm.NewOrm()
	category := &Cms_Knowledge{}
	category.Title = title
	category.Pid = 0
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func SearchKnowledge(pid int) (dataList2 []interface{}, err error) {
	var list []Cms_Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))
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
	var list []Cms_Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))
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
	var list []Cms_Knowledge
	know := Cms_Knowledge{}
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))
	//原生SQL写法
	//o.Raw("select * from cms__knowledge where title = ?", title).QueryRow(&know)
	//orm.Filter
	qs.Filter("title", title).One(&know)
	fmt.Println("GetMoreDirectorys_know", know)
	//查询
	//查询数据
	//if _, err = qs.Filter("pid__exact", 0).Filter("gid__exact", 0).All(&list); err == nil {
	if _, err = qs.Filter("pid", know.Id).All(&list); err == nil {
		fmt.Println("GetMoreDirectorys_list", list)
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

//测试自建结构体插入
func Knowledges(userid interface{}) (dataList []interface{}, err error) {
	var list []Cms_Knowledge
	var data = new(Cms_Knowledge2)
	Uid := userid.(int)
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))

	cond := orm.NewCondition()
	cond1 := cond.And("creater__in", "admin", Uid).Or("status", 1)
	//if _, err = qs.Filter("pid__exact", 0).All(&list); err == nil {            只查询一级目录
	//if _, err = qs.Filter("creater__in", "admin",Uid).All(&list); err == nil { //查询全部
	if _, err = qs.SetCond(cond1).All(&list); err == nil { //查询全部
		for _, v := range list {
			if userid == nil {
				data = &Cms_Knowledge2{v.Id, v.Title, v.Pid, v.Gid, 0, v.UpdateTime, v.Creater, v.Status} //自定义输入内容，前台根据isguanzhu来判定是否显示取消关注，还差一个查询关注表的内容
			} else {
				data = &Cms_Knowledge2{v.Id, v.Title, v.Pid, v.Gid, IsGuanzhu(userid.(int), v.Id), v.UpdateTime, v.Creater, v.Status} //自定义输入内容，前台根据isguanzhu来判定是否显示取消关注，还差一个查询关注表的内容
			}

			dataList = append(dataList, *data)
		}
		fmt.Println("datalist_Knowledge:", dataList)
		return dataList, nil
	}
	return nil, err
}

func InsertKnowledge(title string) error {
	t := time.Now()
	o := orm.NewOrm()
	//category各项属性赋值
	category := &Cms_Knowledge{}
	category.Title = title
	category.Pid = 0
	category.Creater = "admin"
	category.UpdateTime = t.Format("2006-01-02 15:04:05")
	//orm.Insert
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func JumpToKnowledgePage() (dataList []interface{}, err error) {
	var list []Cms_Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))
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
	//原生SQL
	/*res, err := o.Raw("DELETE FROM cms__knowledge  WHERE id = ?", pid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}*/
	//orm.Delete
	num, err := o.Delete(&Cms_Knowledge{Id: pid})
	if err == nil {
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}

func InsertArticle(text, title string, id int) error {
	t := time.Now() //设置当前时间
	o := orm.NewOrm()
	category := &Cms_Article{}
	data := &Cms_Knowledge{}
	data.Title = title
	data.Pid = 0
	data.Creater = "admin"
	data.Gid = id
	data.UpdateTime = t.String()
	dataId, _ := o.Insert(data)
	category.KnowledgeId = int(dataId)
	category.Title = title
	category.Detail = text
	category.Time = t.String() //将时间转换成string类型进行保存
	category.Creater = "admin"
	_, err := o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func GetArticles(id int) (dataList []interface{}, err error) {
	var list []Cms_Article
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Article))
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
	var list []Cms_Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))
	if _, err = qs.Filter("pid__exact", 0).Filter("gid__gt", 0).All(&list); err == nil { //只查询一级目录
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func JumpAllKnowPage(id int) (dataList []interface{}, err error) {
	var list []Cms_Knowledge
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))
	if _, err = qs.Filter("pid__exact", 0).Filter("gid__exact", id).All(&list); err == nil { //只查询一级目录
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

func UserSaveKonwledgeAction(Name string, Uid interface{}) error {
	t := time.Now()
	o := orm.NewOrm()
	//Knowledge各项属性赋值
	Knowledge := &Cms_Knowledge{}
	Knowledge.Title = Name
	Knowledge.Creater = strconv.Itoa(Uid.(int))
	Knowledge.Pid = 0
	Knowledge.Status = 0
	Knowledge.UpdateTime = t.Format("2006-01-02 15:04:05")
	_, err := o.Insert(Knowledge)
	if err != nil {
		return err
	}
	return nil
}

func ChangeKnowledgeStatusAction(Id int, Status int) {
	o := orm.NewOrm()
	Knowledge := Cms_Knowledge{Id: Id}
	if o.Read(&Knowledge) == nil {
		Knowledge.Status = Status
		o.Update(&Knowledge)
	}
}

//保存上传的文件名存入到数据库
func SaveFileName(filename string, id int) error {
	o := orm.NewOrm()
	fmt.Println("SaveFileName_filename=", filename)
	fmt.Println("SaveFileName_id=", id)
	res, err := o.Raw("UPDATE cms__knowledge SET filename = ? WHERE id = ?", filename, id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row update nums:", num)
	}
	return err
}
