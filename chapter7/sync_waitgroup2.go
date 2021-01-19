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
	testFunc := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("%v goroutine start ...\n", id)
		time.Sleep(2)
		fmt.Printf("%v goroutine exit ...\n", id)
	}

	var wg sync.WaitGroup
	const N = 3
	wg.Add(N)
	for i := 0; i < N; i++ {
		go testFunc(&wg, i)
	}

	fmt.Println("Waiting for all goroutine")
	wg.Wait()
	fmt.Println("All goroutines finished!")
}
