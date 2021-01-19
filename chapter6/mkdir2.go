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
	//根据path创建多级子目录，例如dir1/dir2/dir3
	err :=os.MkdirAll("dir1/dir2/dir3", 0777)
	if err != nil {
		fmt.Println(err)
	}
}