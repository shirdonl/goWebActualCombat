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
	// 打开文件
	file, err := os.Open("open.txt")
	if err != nil {
		fmt.Printf("打开文件出错：%v\n", err)
	}
	fmt.Println(file)
	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Printf("关闭文件出错：%v\n", err)
	}
}
