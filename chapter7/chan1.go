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
)
func main() {
	ch := make(chan string) // 构建一个通道
	go func() { // 开启一个并发匿名函数
		fmt.Println("开始协程") // 通过通道通知main的协程
		ch <- "signal"
		fmt.Println("退出协程")
	}()
	fmt.Println("等待协程")
	<-ch // 等待匿名协程
	fmt.Println("完成")
}
