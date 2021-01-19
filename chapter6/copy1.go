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
	"io"
	"os"
)

func main() {
	//先创建一个名为：test_copy1.zip文件
	_, err := os.Create("./test_copy1.zip") // 如果文件已存在，会将文件清空。
	if err != nil {
		fmt.Println(err)
	}
	//打开文件test_copy1.zip，获取文件指针
	srcFile, err := os.Open("./test_copy1.zip")
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer srcFile.Close()

	//打开文件要复制的新文件名test_copy2.zip，获取文件指针
	dstFile, err := os.OpenFile("./test_copy2.zip", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer dstFile.Close()

	//通过Copy方法
	result, err := io.Copy(dstFile, srcFile)

	if err == nil {
		fmt.Println("复制成功，复制的字节数为: ", result)
	}
}
