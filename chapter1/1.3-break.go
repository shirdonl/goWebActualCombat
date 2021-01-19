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
OuterLoop: //外层循环的标签
	for i := 0; i < 2; i++ { //双层循环
		for j := 0; j < 5; j++ {
			switch j { //使用 switch 进行数值分支判断
			case 1:
				fmt.Println(i, j)
				break OuterLoop
			case 2:
				fmt.Println(i, j)
				break OuterLoop //退出 OuterLoop 对应的循环之外
			}
		}
	}
}
