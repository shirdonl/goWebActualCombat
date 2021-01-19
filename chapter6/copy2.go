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
	"log"
	"os"
)

//自定义复制方法
func DoCopy(srcFileName string, dstFileName string) {
	//打开源文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		log.Fatalf("源文件读取失败,err:%v\n", err)
	}
	defer func() {
		err = srcFile.Close()
		if err != nil {
			log.Fatalf("源文件关闭失败,err:%v\n", err)
		}
	}()

	//创建目标文件,稍后会向这个目标文件写入拷贝内容
	distFile, err := os.Create(dstFileName)
	if err != nil {
		log.Fatalf("目标文件创建失败,err:%v\n", err)
	}
	defer func() {
		err = distFile.Close()
		if err != nil {
			log.Fatalf("目标文件关闭失败,err:%v\n", err)
		}
	}()
	//定义指定长度的字节切片,每次最多读取指定长度
	var tmp = make([]byte, 1024*4)
	//循环读取并写入
	for {
		n, err := srcFile.Read(tmp)
		n, _ = distFile.Write(tmp[:n])
		if err != nil {
			if err == io.EOF {
				return
			} else {
				log.Fatalf("拷贝过程中发生错误,错误err:%v\n", err)
			}
		}
	}
}

func main() {
	//先创建一个.zip文件
	_, err := os.Create("./test.zip") // 如果文件已存在，会将文件清空。
	if err != nil {
		fmt.Println(err)
	}
	//复制一个名为test2.zip的文件
	DoCopy("./test.zip", "./test2.zip")
}
