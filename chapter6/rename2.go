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
	//创建一个名为“test_rename.txt”的空文件
	_, err := os.Create("./test_rename.txt") // 如果文件已存在，会将文件清空。
	if err != nil {
		fmt.Println(err)
	}
	//创建一个名为“test_rename”的目录，perm权限为0777
	err = os.Mkdir("test_rename", 0777)
	//将test_rename.txt移动到test_rename目录，并将名字重命名为test_rename_new.txt
	err = os.Rename("./test_rename.txt", "./test_rename/test_rename_new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}