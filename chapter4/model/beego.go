//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package model

import "github.com/astaxie/beego/orm"

type BeegoUser struct {
	Id   int //默认主健为id
	Name  string
	Phone string
}

func init(){
	orm.RegisterModel(new(BeegoUser))
}
