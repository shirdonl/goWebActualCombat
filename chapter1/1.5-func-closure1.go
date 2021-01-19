//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import "fmt"

func main() {
	// 定义匿名函数并赋值给f变量
	f := func(data int) {
		fmt.Println("hi, this is a closure", data)
	}
	// 此时f变量的类型是func(), 可以直接调用
	f(6)

	//直接声明并调用
	func(data int) {
		fmt.Println("hi, this is a closure, directly", data)
	}(8)
}
