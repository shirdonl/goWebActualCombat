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
	// 使用命令行提高拷贝的复用性
	args := os.Args
	if args == nil || len(args) != 3 {
		fmt.Println("useage : go filename.go src File dstFile")
		return
	}
	srcPath := args[1]
	dstPath := args[2]
	fmt.Printf("srcPath = %s, dstPath = %s\r\n", srcPath, dstPath)
	if srcPath == dstPath {
		fmt.Println("源文件和目标文件不能重名")
	}
	//执行复制
	srcFile, err1 := os.Open(srcPath)
	// 关闭文件
	defer srcFile.Close()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	dstFile, err2 := os.Create(dstPath)
	defer dstFile.Close()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	readBuf := make([]byte, 1024)
	for {
		//读取文件
		n, err := srcFile.Read(readBuf) //每次文件读取字节的长度
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}
		if n == 0 {
			fmt.Println("文件处理完毕")
		}
		//写入目的文件
		writeBuf := readBuf[:n]
		dstFile.Write(writeBuf)
	}
}
