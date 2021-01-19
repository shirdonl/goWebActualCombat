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
func IntegerGenerator() chan int {
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		for i := 2; ; i++ {
			ch <- i  // 直到通道索要数据，才把i添加进通道
		}
	}()

	return ch
}

func Filter(in chan int, number int) chan int {
	// 输入一个整数队列，筛出是number倍数的, 不是number的倍数的放入输出队列
	// in:  输入队列
	out := make(chan int)

	go func() {
		for {
			i := <-in // 从输入中取一个

			if i%number != 0 {
				out <- i // 放入输出通道
			}
		}
	}()

	return out
}

func main() {
	const max = 100               // 找出100以内的所有素数
	numbers := IntegerGenerator() // 初始化一个整数生成器
	number := <-numbers           // 从生成器中抓一个整数(2), 作为初始化整数

	for number <= max { // number作为筛子，当筛子超过max的时候结束筛选
		fmt.Println(number)               // 打印素数, 筛子即一个素数
		numbers = Filter(numbers, number) //筛掉number的倍数
		number = <-numbers                // 更新筛子
	}
}
