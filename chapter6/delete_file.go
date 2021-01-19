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
	"os"
)

func main() {
	err := os.Remove("./file1.txt")
	if err != nil {
		fmt.Printf("remove ./file1.txt err : %v\n", err)
	}
	err = os.RemoveAll("./file2.txt")
	if err != nil {
		fmt.Printf("remove all ./file2.txt err : %v\n", err)
	}
}
