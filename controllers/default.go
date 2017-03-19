package controllers

import (
	"github.com/astaxie/beego"
	"myGoApp/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName="index.html"
	c.Data["IsIndex"]=true
	c.Data["IsLogin"]=checkAccout(c.Ctx)
	var err error
	c.Data["Topics"],err=models.GetAllTopics(true)
	if err!=nil{
		beego.Error(err)
	}
}
