//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

func main() {
	var var1, var2 interface{}
	println(var1 == nil, var1 == var2)
	var1, var2 = 66, 88
	println(var1 == var2)
	var1, var2 = map[string]string{}, map[string]string{}
	println(var1 == var2)
}
