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
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)
	// 写的时候啥也不能干
	go Writing(1)
	go Read(2)
	go Writing(3)
	time.Sleep(2 * time.Second)
}
func Read(i int) {
	println(i, "reading start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	println(i, "reading over")
}
func Writing(i int) {
	println(i, "writing start")
	m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()
	println(i, "writing over")
}
