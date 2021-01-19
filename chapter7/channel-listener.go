//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebFromIntroductionToMasterb
// 仓库地址：https://github.com/shirdonl/goWebFromIntroductionToMasterb
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
)

func foo(i int) chan int {
	ch := make(chan int)
	go func() { ch <- i }()
	return ch
}

func main() {
	ch1, ch2, ch3 := foo(3), foo(6), foo(9)

	ch := make(chan int)

	// 开一个goroutine监视各个通道数据输出并收集数据到通道ch
	go func() {
		for {
			// 监视ch1, ch2, ch3的流出，并全部流入通道ch
			select {
			case v1 := <-ch1:
				ch <- v1
			case v2 := <-ch2:
				ch <- v2
			case v3 := <-ch3:
				ch <- v3
			}

			//timeout 是一个计时通道, 如果达到时间了，就会发一个信号出来
			//timeout := time.After(1 * time.Second)
			//for isTimeout := false; !isTimeout; {
			//	select { // 监视通道ch1, ch2, ch3, timeout通道的数据流出
			//	case v1 := <-ch1:
			//		fmt.Printf("received %d from ch1", v1)
			//	case v2 := <-ch2:
			//		fmt.Printf("received %d from ch2", v2)
			//	case v3 := <-ch3:
			//		fmt.Printf("received %d from ch3", v3)
			//	case <-timeout:
			//		isTimeout = true // 超时
			//	}
			//}
		}
	}()

	// 阻塞主线，取出通道ch的数据
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}


}
