package models

import (
	"os"
	"time"

	"github.com/astaxie/beego/orm"
	_"github.com/mattn/go-sqlite3"   //sqllite3的数据库驱动
	"github.com/Unknwon/com"

	"path"
	"strconv"
)
/**
*beego的mvc框架的model层，并利用提供的orm工具，目前支持mysql，sqlite ,pg数据库
*(1)定义数据库中表的结构,设置索引，一对一，一对多等对应关系
*
 */
const (
	_DB_NAME  ="data/beeblog.db"  //定义数据库文件的位置，设置为常量
	_SQLITE3_DRIVER ="sqlite3"    //定义数据库驱动
)


type Category struct {
	Id int64
	Title string
	Created time.Time 	`orm:"index"`   //设置索引
	Views int64		`orm:"index"`
	TopicTime time.Time	`orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}

type Topic struct {
	Id int64
	Uid int64
	Title string
	Category string
	Content string		`orm:"size(5000)"`
	Attachment string
	Created time.Time	`orm:"index"`
	Updated time.Time	`orm:"index"`
	Views int64		`orm:"index"`
	Author string
	ReplyTime time.Time	`orm:"index"`
	ReplyCount int64
	ReplyLastUserId int64
}

/*
注册数据库的相关信息
 */

func RegisterDB(){
    	//检查目录是否存在，需要用到path包
    	if !com.IsExist(_DB_NAME){
	    os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
	    os.Create(_DB_NAME)
    	}
	/*
	注册模型，这里看到将category，topic的注册和对应的方法写到了一个文件，最佳设计应该是分开单独在各自的init函数中进行注册
	 */
	orm.RegisterModel(new(Category),new(Topic))

	/*
	注册数据库驱动 1.6版本之前时DR.Sqlite
	参数1：驱动名称
	参数2：数据库类型
	 */
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)

	/*
	注册数据库,orm必须注册一个别名为default的数据库，作为默认使用
	参数1：数据库的别名
	参数2：驱动名称
	参数3：对应的链接字符串
	参数4（可选）：设置最大空闲连接 也可以单独设置 orm.setMaxIdleConns("default",30)
	参数5（可选）：设置最大数据库连接 也可以单独设置 orm.setMaxOpenConns("default",30)
	 */
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)
}

/**
下面定义了一系列操作数据库对应的方法，数据库的curd等操作
 */

//添加数据到category
func AddCategory(name string) error  {
	o:=orm.NewOrm() //创建一个ormer ，如果只是用到一个数据库的话，可以定义一个全局的ormer
	/*
	这是一种操作方式，还可以直接写sql语句来操作 o.Raw("sql",param1,param2,...)
	 */
	cate:=&Category{Title:name} //依据传入的name构造一个category对象
	qs:=o.QueryTable("category") //传入表名
	err:=qs.Filter("title",name).One(cate) //过滤器
	if err==nil{
		return err
	}
	/*
	执行插入操作，第一个返回自增的id键值，第二个返回异常
	 */
	_,err =o.Insert(cate)
	if err!=nil{
		return err
	}
	return nil
}

//从category中查询数据
func GetAllCategories() ([]*Category,error) {
	o:=orm.NewOrm()
	cates:=make([]*Category,0) //返回值
	qs:=o.QueryTable("category")
	_,err:=qs.All(&cates)
	return cates,err
}

//从category中删除数据
func DelCategory(id string) error  {
	cid,err:=strconv.ParseInt(id,10,64)
	if err!=nil{
		return err
	}

	o:=orm.NewOrm()

	cate:=&Category{Id:cid}
	_,err=o.Delete(cate)
	return err
}

//添加文章到topic
func AddTopic(title,content,category string) error {
	o:=orm.NewOrm()

	topic:=&Topic{
		Title:title,
		Content:content,
		Category:category,
		Created:time.Now(),
		Updated:time.Now(),
	}

	_,err:=o.Insert(topic)
	return err
}

//
func GetAllTopics(isDesc bool) ([]*Topic,error){
	o:=orm.NewOrm()
	topics:=make([]*Topic,0)
	qs:=o.QueryTable("topic")
	var err error
	if isDesc{
		_,err=qs.OrderBy("-Created").All(&topics)
	}else{
		_,err=qs.All(&topics)
	}

	return topics,err
}

//
func GetTopic(tid string) (*Topic,error){
	tidNum,err:=strconv.ParseInt(tid,10,64)
	if err!=nil{
		return nil,err
	}

	o:=orm.NewOrm()

	topic:=new(Topic)

	qs:=o.QueryTable("topic")
	err=qs.Filter("id",tidNum).One(topic)

	if err!=nil{
		return nil,err
	}

	topic.Views++

	_,err=o.Update(topic)
	return topic,err
}

func ModifyTopic(tid ,title,content,category string) error{
	tidNum,err:=strconv.ParseInt(tid,10,64)
	if err!=nil{
		return err
	}

	o:=orm.NewOrm()

	topic:=&Topic{
		Id:tidNum,
	}

	if o.Read(topic)==nil{
		topic.Title=title
		topic.Content=content
		topic.Category=category
		topic.Updated=time.Now()
		o.Update(topic)
	}
	return  nil
}

func DeleteTopic(tid string)  error {
	cid,err:=strconv.ParseInt(tid,10,64)
	if err!=nil{
		return err
	}

	o:=orm.NewOrm()

	topic:=&Topic{Id:cid}
	_,err=o.Delete(topic)
	return err
}