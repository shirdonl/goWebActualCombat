package main

import (
	"fmt"
)

func main() {

	ch, quit := make(chan int), make(chan int)

	go func() {
		ch <- 8   // 添加数据
		quit <- 1 // 发送完成信号
	}()

	for isQuit := false; !isQuit; {
		// 监视通道ch的数据流出
		select {
		case v := <-ch:
			fmt.Printf("received %d from ch", v)
		case <-quit:
			isQuit = true // quit通道有输出，关闭for循环
		}
	}
}
