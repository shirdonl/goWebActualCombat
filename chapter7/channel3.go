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
	ch := make(chan int, 2)
	// 发送
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			ch <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch)
	}()

	// 接收
	for {
		elem, ok := <-ch
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}
