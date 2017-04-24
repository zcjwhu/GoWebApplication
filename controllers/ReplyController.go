package controllers

import "github.com/astaxie/beego"
import "myGoApp/models"

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add(){
	tid:=this.Input().Get("tid")
	err:=models.AddReply(tid,this.Input().Get("nickname"),this.Input().Get("content"))
	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid,302)
}

func (this *ReplyController) Delete(){
	err:=models.DeleteReply(this.Input().Get("rid"))
	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+this.Input().Get("tid"),302)
}