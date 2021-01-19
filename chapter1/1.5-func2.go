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
	array := []int{6, 8, 10} //定义局部变量
	var ret int
	ret = min(array) //调用函数并返回最小值
	fmt.Printf("最小值是 : %d\n", ret)
}

func min(arr []int) (min int) { //获取整型数组中的最小值
	min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}
