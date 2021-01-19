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

	//var (
	//	age     int = 28
	//	name    string = "shirdon"
	//	balance float32 = 999999.99
	//)

	//var (
	//	age int
	//	name string
	//	balance float32
	//)

	var language string = "Go"
	//var language = "Go"
	//language := "Go"

	//var age, name, balance = 30, "liao", 99.99
	age, name, balance := 30, "shirdon", 999999.99

	d, c := "D", "C"
	c, d = d, c

	println(c)
	println(d)

	println(age)
	println(name)
	println(balance)
	println(language)
}
