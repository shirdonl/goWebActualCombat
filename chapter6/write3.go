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
	"os"
)

func main() {
	file, err := os.Create("WriteString.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Go Web编程实战派从入门到精通")
}
