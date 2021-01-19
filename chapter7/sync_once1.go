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
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("test only once，这里只打印一次！")//打印
	}
	done := make(chan bool)
	for i := 0; i < 6; i++ {
		go func() {
			once.Do(onceBody)//确保只背执行一次
			done <- true
		}()
	}
	for i := 0; i < 6; i++ {
		<-done
	}
}
