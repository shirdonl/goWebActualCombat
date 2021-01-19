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
	"math/rand"
	"time"
)

//这里可以是比较耗时的事情，比如计算
func doCompute(x int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) //模拟计算
	return 1 + x                                                // 假如1 + x是一个很费时的计算
}

// 每个分支开出一个goroutine做计算，并把计算结果发送到各自通道
func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- doCompute(x)
	}()
	return ch
}

func Recombination(chs ... chan int) chan int {
	ch := make(chan int)

	for _, c := range chs {
		// 注意此处明确传值
		go func(c chan int) { ch <- <-c }(c) // 复合
	}

	return ch
}

func Recombination(branches ... chan int) chan int {
	ch := make(chan int)

	//select会尝试着依次取出各个通道的数据
	go func() {
		for i := 0; i < len(branches); i++ {
			select {
			case v1 := <-branches[i]:
				ch <- v1
			}
		}
	}()

	return ch
}

func main() {
	//返回复合结果
	result := Recombination(branch(10), branch(20), branch(30))

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
