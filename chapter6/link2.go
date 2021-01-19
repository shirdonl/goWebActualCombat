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
	// 创建文件
	//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666//-rw-r--r--
	fp, err := os.Create("./link2.txt") // 如果文件已存在，会将文件清空。
	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。
	if err != nil {
		fmt.Println("文件创建失败。")
	}
	err = os.Symlink("link2.txt", "link3.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
}
