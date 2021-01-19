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
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "http://www.baidu.com:80")
	if err != nil {
		//处理错误
	}
	//读写操作

	conn, err := net.DialTimeout("tcp", ":8085", 2 * time.Second)
	if err != nil {
		//处理错误
	}
	//读写操作
}
