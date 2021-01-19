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
	// 读写方式打开文件
	fp, err := os.OpenFile("./open.txt", os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("文件打开失败。")
		return
	}

	// defer延迟调用
	defer fp.Close()  //关闭文件，释放资源。
}
