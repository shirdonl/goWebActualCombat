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
	// 声明局部变量
	var local1, local2, local3 int

	//初始化参数
	local1 = 8
	local2 = 10
	local3 = local1 * local2

	fmt.Printf(" local1 = %d, local2 = %d and local3 = %d\n", local1, local2, local3)
}
