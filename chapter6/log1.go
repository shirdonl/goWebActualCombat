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
	"log"
)

func main() {
	no := []int{6, 8}
	log.Print("Print NO. ", no, "\n")
	log.Println("Println NO.", no)
	log.Printf("Printf NO. with item [%d,%d]\n", no[0], no[1])
}
