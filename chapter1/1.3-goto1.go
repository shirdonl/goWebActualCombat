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
	//var isBreak bool
	//for x := 0; x < 20; x++ {// 外循环
	//	for y := 0; y < 20; y++ {// 内循环
	//		if y == 2 {// 满足某个条件时, 退出循环
	//			isBreak = true// 设置退出标记
	//			break// 退出本次循环
	//		}
	//	}
	//	if isBreak {// 根据标记, 还需要退出一次循环
	//		break
	//	}
	//}
	//fmt.Println("over")

	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			if y == 2 {
				goto breakTag // 跳转到标签
			}
		}
	}
	return // 手动返回, 避免执行进入标签
breakTag: // 标签
	fmt.Println("done")
}
