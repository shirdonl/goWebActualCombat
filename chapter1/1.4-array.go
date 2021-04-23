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
	var arr [6]int // arr 是一个长度为 6 的数组
	var i, j int
	for i = 0; i < 6; i++ { // 为数组 arr 初始化元素
		arr[i] = i + 66 // 设置元素为 i + 66
	}
	for j = 0; j < 6; j++ { // 输出每个数组元素的值
		fmt.Printf("Array[%d] = %d\n", j, arr[j])
	}
}
