package controllers

import "github.com/astaxie/beego"

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get(){
	c.Ctx.WriteString("this is my hello url test")
}
