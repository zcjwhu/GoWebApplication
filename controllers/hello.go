package controllers

import "github.com/astaxie/beego"

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get(){
	c.Data["Website"] = "hello world"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	//模板中使用条件语句
	c.Data["TrueCond"]=true
	c.Data["FalseCond"]=false

	//结构类型的数据
	type u struct {
		Name string
		Age int
		Sex string
	}

	user:=&u{
		Name:"Joe",
		Age:25,
		Sex:"男",
	}
	c.Data["User"]=user

	//数组数据的传递，模板中使用循环结构
	nums:=[]int{1,2,3,4,5,6,7,8,9,0}
	c.Data["Nums"]=nums

	//模板中获取变量
	c.Data["variable"]="zcj"

	//模板函数的使用，字符串到html
	c.Data["html"]="<div>hello beego</div>"
}
