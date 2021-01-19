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
	"os"
)

func main() {
	//先创建多级子目录
	os.MkdirAll("test1/test2/test3", 0777)
	//删除test1目录及其子目录
	err := os.RemoveAll("test1")
	if err != nil {
		log.Fatal(err)
	}
}