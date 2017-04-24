package main

import (
	"myGoApp/models"    //model数据结构模块
	_ "myGoApp/routers"  //路由映射模块    前面加上一个下划线表示的是只调用这个模块里面的init()函数去执行
	"github.com/astaxie/beego"   //beego包  下面这两个包都会从本机gopath里面按照路径去寻找
	"github.com/astaxie/beego/orm" //orm包
)

func init()  {
	/*
        这一步对所有注册的模型进行注册，实际上可以对每个模型在各自的init方法里面注册，然后使用_,将包引入，会自动执行对应的init方法
	 */
	models.RegisterDB()
}

func main() {
	orm.Debug=true  //开发模式下设置为true,用于显示开发过程中的一些数据
	orm.RunSyncdb("default",false,true) //false表示每次会强调删除然后重建 ，true表示是否打印相关信息
	beego.Run()
}

