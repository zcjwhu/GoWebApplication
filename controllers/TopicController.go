package controllers

import (
	"github.com/astaxie/beego"
	"myGoApp/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get()  {
	this.Data["IsTopic"]=true
	this.TplName="topic.html"

	var err error
	this.Data["Topics"],err=models.GetAllTopics(false)
	if err!=nil{
		beego.Error(err)
	}
}

func (this *TopicController) Add(){
	this.TplName="topic_add.html"
}

func (this *TopicController) Post(){
         if !checkAccout(this.Ctx){
		this.Redirect("/login",302)
		return
	}

	title:=this.Input().Get("title")
	content:=this.Input().Get("content")
	category:=this.Input().Get("category")

	var err error
	var tid string
	tid=this.Input().Get("tid")
	if len(tid) ==0{
		err= models.AddTopic(title,content,category)
	}else{
		err=models.ModifyTopic(tid,title,content,category)
	}

	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/topic",302)
}

func (this *TopicController) View(){
	this.TplName="topic_view.html"
	topic,err:=models.GetTopic(this.Ctx.Input.Param("0"))
	if err!=nil{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}

	this.Data["Topic"]=topic
	this.Data["Tid"]=this.Ctx.Input.Param("0")
	replies,err:=models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err!=nil{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}
	this.Data["Replies"]=replies
}

func (this *TopicController) Modify(){
	this.TplName="topic_modify.html"
	tid:=this.Input().Get("tid")
	topic,err:=models.GetTopic(tid)
	if err!=nil{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}

	this.Data["Topic"]=topic
	this.Data["Tid"]=tid
}

func (this *TopicController) Delete()  {
	if !checkAccout(this.Ctx){
		this.Redirect("/login",302)
		return
	}

	err:=models.DeleteTopic(this.Ctx.Input.Param("0"))
	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/",302)
}