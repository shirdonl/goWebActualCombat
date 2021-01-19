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

//生成自增的整数
func IntegerGenerator() chan int{
	var ch chan int = make(chan int)

	// 开启 goroutine
	go func() {
		for i := 0; ; i++ {
			ch <- i  // 直到通道索要数据，才把i添加进信道
		}
	}()

	return ch
}

func main() {

	generator := IntegerGenerator()

	for i:=0; i < 100; i++ {  //生成100个自增的整数
		fmt.Println(<-generator)
	}
}
