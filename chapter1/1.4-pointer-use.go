//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import (
	"fmt"
)

func main() {
	var score int = 100
	var name string = "Barry"
	// 使用 fmt.Printf 的动词%p输出 score 和 name 变量取地址后的指针值
	fmt.Printf("%p %p", &score, &name)
}
