package main

import (
	"myGoApp/models"    //model数据结构模块
	_ "myGoApp/routers"  //路由映射模块
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm" //orm包
)

func init()  {
	models.RegisterDB()
}

func main() {
	orm.Debug=true  //开发模式下设置为true,用于显示开发过程中的一些数据
	orm.RunSyncdb("default",false,true) //false表示每次会强调删除然后重建 ，true表示是否打印相关信息
	beego.Run()
}

