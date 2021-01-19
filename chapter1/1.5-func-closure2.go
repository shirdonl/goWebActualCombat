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

// 遍历切片中每个元素，通过给定的函数访问元素
func visitPrint(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}

func main() {
	sli := []int{1, 6, 8}
	// 使用匿名函数打印切片的内容
	visitPrint(sli, func(value int) {
		fmt.Println(value)
	})
}
