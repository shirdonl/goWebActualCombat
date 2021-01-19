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
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")
}
