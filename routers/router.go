package routers

import (
	"myGoApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//注册路由
      	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello",&controllers.HelloController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/category",&controllers.CategoryController{})
	beego.Router("/reply",&controllers.ReplyController{})
	beego.Router("/reply/add",&controllers.ReplyController{},"post:Add")
	beego.Router("/reply/delete",&controllers.ReplyController{},"get:Delete")
	beego.Router("/topic",&controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{}) //自动路由，需要异controller来结尾



}
