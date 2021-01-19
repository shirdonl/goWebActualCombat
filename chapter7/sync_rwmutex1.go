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
	"sync"
)

func main() {
	var num int
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	go func() {
		defer wg.Done()
		mutex.Lock()
		num = 4
		mutex.Unlock()
		fmt.Printf("goroutine one, num=%d\n", num)
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		num = 5
		mutex.Unlock()
		fmt.Printf("goroutine two, num=%d\n", num)
	}()

	wg.Add(2)
	wg.Wait()
	fmt.Println("end main goroutine")
}
