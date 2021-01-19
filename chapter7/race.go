//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import "fmt"

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["a"] = "one" // 第一个冲突访问.
		c <- true
	}()
	m["b"] = "two" // 第一个冲突访问
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
