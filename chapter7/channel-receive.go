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
	"time"
)

func main() {
	ch := make(chan int) // 构建一个通道
	go func() {          // 开启一个并发匿名函数
		for i := 6; i <= 8; i++ { // 从6循环到8
			ch <- i                 // 发送6到8之间的数值
			time.Sleep(time.Second) // 每次发送完时等待
		}
	}()

	for receive := range ch { // 遍历接收通道数据
		fmt.Println(receive) // 打印通道数据
		if receive == 8 {    // 当遇到数据8时, 退出接收循环
			break
		}
	}
}
