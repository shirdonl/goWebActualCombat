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
	"log"
	"os"
)

func main() {
	//创建一个名为“dir_name1”的目录，perm权限为0777
	err := os.Mkdir("dir_name1", 0777)
	if err != nil {
		fmt.Println(err)
	}
	oldName := "dir_name1"
	newName := "dir_name2"
	//将dir_name1重命名为dir_name2
	err = os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}