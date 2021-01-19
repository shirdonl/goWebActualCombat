package main

import "fmt"

func randGenerator() chan int {
	ch := make(chan int)

	go func() {
		for {
			//select会尝试执行各个case, 如果都可以执行，那么随机选一个执行
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	return ch
}

func main() {
	//初始化一个随机生成器
	generator := randGenerator()

	//测试，打印10个随机数
	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}
}
