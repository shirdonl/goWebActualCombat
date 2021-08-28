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

func Sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum // 把 sum 发送到通道 ch
}

func main() {
	s := []int{6, 7, 8, -9, 1, 8}
	ch := make(chan int, 6)

	go Sum(s[:len(s)/2], ch)
	a := <-ch
	go Sum(s[len(s)/2:], ch)
	b := <-ch

	//fmt.Println(<-ch)
	//fmt.Println(s[:len(s)/2])
	//a, b := <-ch, <-ch // 从通道 ch 中接收
	fmt.Println(a, b, a+b)
}
