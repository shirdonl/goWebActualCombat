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

// 交换函数
func exchange(c, d *int) {
	t := *c // 取c指针的值, 赋给临时变量t
	*c = *d // 取d指针的值, 赋给c指针指向的变量
	*d = t // 将c指针的值赋给d指针指向的变量
}
func main() {
	a, b := 6, 8      // 准备两个变量, 赋值6和8
	exchange(&a, &b)  // 交换变量值
	fmt.Println(a, b) // 输出变量值
}

