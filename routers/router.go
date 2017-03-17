package routers

import (
	"myGoApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
      	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello",&controllers.HelloController{})
}
