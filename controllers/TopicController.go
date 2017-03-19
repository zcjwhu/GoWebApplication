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

	var err error
	err= models.AddTopic(title,content)

	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/topic",302)
}