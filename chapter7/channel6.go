//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebFromIntroductionToMasterb
// 仓库地址：https://github.com/shirdonl/goWebFromIntroductionToMasterb
//++++++++++++++++++++++++++++++++++++++++

package main

import "fmt"

func main() {

	var a, b int
	go func() {
		a = 1
		fmt.Println("a:", b, " ")
	}()
	go func() {
		b = 1
		fmt.Println("b:", a, " ")
	}()
}
