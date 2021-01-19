package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)
	// 多个同时读
	go Reading(1)
	go Reading(2)
	time.Sleep(2 * time.Second)
}
func Reading(i int) {
	println(i, "reading start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	println(i, "reading over")
}
