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
	x, y := 6, 8
	defer func(a int) {
		fmt.Println("defer x, y = ", a, y) //y为闭包引用
	}(x)
	x += 10
	y += 100
	fmt.Println(x, y)
}
