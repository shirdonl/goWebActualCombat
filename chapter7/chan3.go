package main

import "fmt"

func main() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道，缓冲区大小为3
	ch := make(chan int, 3)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 6
	ch <- 7
	ch <- 8
	// 获取这三个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
