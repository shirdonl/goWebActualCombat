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
	//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666
	fp, err := os.Create("./demo.txt") // 如果文件已存在，会将文件清空。
	fmt.Println(fp, err)               // &{0xc000054180} <nil>
	fmt.Printf("%T", fp)               // *os.File 文件指针类型

	if err != nil {
		fmt.Println("文件创建失败。")
		//创建文件失败的原因有：
		//1、路径不存在  2、权限不足  3、打开文件数量超过上限  4、磁盘空间不足等
		return
	}

	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。
}
