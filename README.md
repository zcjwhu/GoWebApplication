Go语言结合Beego框架开发Web应用
====
工具环境：
---
###1.Go语言<br>
1.下载并安装Go
2.配置GOROOT(指go语言的安装目录/usr/local/go/bin)和GOPATH(指go语言的工作目录/mygopath),go工程的工作目录一般分为src、pkg、bin,使用go get获取的包都会放到src目录下
###2.Beego包
使用go语言开发的一个Web框架 获取方式go get github.com/astaxie/beego<br> 
###3.Bee包
用于构建项目、编译、运行项目的命令行工具 获取方式go get github.com/beego/bee<br>
###4.orm包
它是beego附带的做orm的一个包<br>
###5.sqlite包
sqllite数据库 获取方式go get github.com/mattn/go-sqlite3
