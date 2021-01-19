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
	"path/filepath"
)

func scan(path string, f os.FileInfo, err error) error {
	fmt.Printf("Scaned: %s\n", path)
	return nil
}

func main() {
	//根据path创建多级子目录，例如dir1/dir2/dir3
	err :=os.MkdirAll("test_walk/dir2/dir3", 0777)
	if err != nil {
		fmt.Println(err)
	}
	root := `./test_walk`
	err = filepath.Walk(root, scan)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}