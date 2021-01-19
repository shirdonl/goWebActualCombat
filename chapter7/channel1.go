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
	"runtime"
	"sync"
)

var Counter int

func Count(lock *sync.Mutex) {
	lock.Lock() // 上锁
	Counter++
	fmt.Println("Counter =", Counter)
	lock.Unlock() // 解锁
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 6; i++ {
		go Count(lock)
	}
	for {
		lock.Lock() // 上锁
		c := Counter
		lock.Unlock() // 解锁
		runtime.Gosched() // 出让时间片
		if c >= 6 {
			break
		}
	}
}
