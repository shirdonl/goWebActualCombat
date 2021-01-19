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
	"time"
)

func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	fmt.Println("Locked")
	mutex.Lock()
	for i := 1; i <= 5; i++ {
		wait.Add(1)
		go func(i int) {
			fmt.Println("Not lock:", i)
			mutex.Lock()
			fmt.Println("Lock:", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock:", i)
			mutex.Unlock()
			defer wait.Done()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()
	wait.Wait()
}