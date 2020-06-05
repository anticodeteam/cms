package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	Filename   string
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
type Cms_Comments struct {
	Id          int
	Knowledgeid int
	Name        string
	Comments    string
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
	var user = Cms_User{}
	var list []Cms_Knowledge
	var data = new(Cms_Knowledge2)
	Uid := userid.(int) //处理输入参数userid的数据类型
	user.Uid = Uid      //将处理后的userid赋值给user对象
	o := orm.NewOrm()
	err_auth := o.Read(&user) //根据userid去判断当前用户的权限
	qs := o.QueryTable(new(Cms_Knowledge))
	cond := orm.NewCondition()
	cond1 := cond
	if err_auth == nil {
		if user.Auth == "1" { //普通用户
			cond1 = cond.And("creater__in", "admin", Uid).Or("status__exact", 3)
		} else if user.Auth == "2" { //管理员
			cond1 = cond.And("creater__exact", "admin").Or("status__in", 1, 2)
		}
	}
	if _, err = qs.SetCond(cond1).All(&list); err == nil { //查询全部
		for _, v := range list {
			if userid == nil {
				data = &Cms_Knowledge2{v.Id, v.Title, v.Pid, v.Gid, 0, v.UpdateTime, v.Creater, v.Status} //自定义输入内容，前台根据isguanzhu来判定是否显示取消关注，还差一个查询关注表的内容
			} else {
				data = &Cms_Knowledge2{v.Id, v.Title, v.Pid, v.Gid, IsGuanzhu(userid.(int), v.Id), v.UpdateTime, v.Creater, v.Status} //自定义输入内容，前台根据isguanzhu来判定是否显示取消关注，还差一个查询关注表的内容
			}
			dataList = append(dataList, *data)
		}
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

func InsertArticle(text, title string, id int, Uid interface{}) error {
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
	category.Time = t.Format("2006-01-02 15:04:05") //将时间转换成string类型进行保存
	category.Creater = strconv.Itoa(Uid.(int))
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

//获取表中评论
func GetCommentList(id int) (dataList []interface{}, err error) {
	fmt.Println("id=====", id)
	var list []Cms_Comments
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Comments))
	//fmt.Println("list=",qs.Filter("knowledgeid",id))
	//res, err := o.Raw("SELECT * FROM cms_comments WHERE knowledgeid = ?", id).Exec()
	//fmt.Println("res=",res)
	if _, err = qs.Filter("knowledgeid", id).All(&list); err == nil {
		fmt.Println("list", list)
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
		fmt.Println(dataList)
	}
	return nil, err
}

//保存评论
func SaveComments(id interface{}, name string, comment string) {
	o := orm.NewOrm()
	comments := &Cms_Comments{}
	comments.Knowledgeid = id.(int)
	comments.Name = name
	comments.Comments = comment
	id, err := o.Insert(comments)
	if err == nil {
		fmt.Println(id)
	}

}

//删除评论
func DeleteComment(name string) {
	o := orm.NewOrm()
	fmt.Println("Name=", name)
	num, err := o.Delete(&Cms_Comments{Name: name})
	if err != nil {
		beego.Info("删除失败", err)
		logs.Info("删除失败", err)
		return
	}
	beego.Info("删除成功，一共删除了：", num, "条")
}

//
func ChangeThisStatusAction(Id int, status int) {
	o := orm.NewOrm()
	know := Cms_Knowledge{Id: Id}
	if o.Read(&know) == nil {
		know.Status = status
		if num, err := o.Update(&know); err == nil {
			fmt.Println(num)
		}
	}

}

//
func AddLevel2MenuAction(Id int, title string, Uid interface{}) {
	o := orm.NewOrm()
	t := time.Now()
	var know Cms_Knowledge
	know.Pid = Id
	know.Title = title
	know.Status = 0
	know.Creater = strconv.Itoa(Uid.(int))
	know.Gid = 0
	know.UpdateTime = t.Format("2006-01-02 15:04:05")

	id, err := o.Insert(&know)
	if err == nil {
		fmt.Println(id)
	}
}

//
func DeleteKnowAction(id int) {
	var list []Cms_Knowledge
	var knowids []int
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cms_Knowledge))
	qs.Filter("pid", id).All(&list)
	fmt.Println("DeleteKnowAction->list:", list)
	if len(list) > 1 {
		for _, v := range list {
			knowids = append(knowids, v.Id)
		}
		for _, val := range knowids {
			o.Delete(&Cms_Knowledge{Id: val})
		}
		o.Delete(&Cms_Knowledge{Id: id})
	} else {
		if num, err := o.Delete(&Cms_Knowledge{Id: id}); err == nil {
			fmt.Println(num)
		}
	}
}
