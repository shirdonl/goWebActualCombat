//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

//CREATE TABLE `beego_user` (
// `id` bigint(20) NOT NULL AUTO_INCREMENT,
// `name` varchar(20) DEFAULT '',
// `phone` varchar(20) DEFAULT '',
// PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4

package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //数据库类型设计
	orm.RegisterDataBase("default", "mysql",
		"root:a123456@tcp(127.0.0.1:3306)/chapter4?charset=utf8")
	//需要在init中注册定义的model
	orm.RegisterModel(new(BeegoUser))
}

type BeegoUser struct {
	Id   int //默认主健为id
	Name  string
	Phone string
}

func main() {
	o := orm.NewOrm()
	user := new(BeegoUser)
	user.Name = "Shirdon"
	user.Phone = "18888888888"
	fmt.Println(o.Insert(user))
}