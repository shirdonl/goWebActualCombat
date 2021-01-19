package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件
	//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666//-rw-r--r--
	fp, err := os.Create("./chmod1.txt") // 如果文件已存在，会将文件清空。
	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。
	if err != nil {
		fmt.Println("文件创建失败。")
	}
	fileInfo, err := os.Stat("./chmod1.txt")
	fileMode := fileInfo.Mode()
	fmt.Println(fileMode)
	os.Chmod("./chmod1.txt", 0777)//通过chmod重新赋权限//-rwxrwxrwx
	fileInfo, err =os.Stat("./chmod1.txt")
	fileMode = fileInfo.Mode()
	fmt.Println(fileMode)
}
