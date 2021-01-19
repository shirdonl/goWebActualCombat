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
	"time"
)

func main()  {
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir , 777)
	if err!=nil{
		fmt.Println(err)
	}
}
