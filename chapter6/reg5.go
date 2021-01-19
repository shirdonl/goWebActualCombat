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
	"strings"
)

func main() {
	value := "i'm gopher"
	result := strings.SplitAfter(value, "")
	for i := range (result) {
		// 获取字母
		letter := result[i]
		fmt.Println(letter)
	}
}
