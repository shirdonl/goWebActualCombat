//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

//接口1
type Interface1 interface {
	Write(p []byte) (n int, err error)
}

//接口2
type Interface2 interface {
	Close() error
}

//接口组合
type InterfaceCombine interface {
	Interface1
	Interface2
}
