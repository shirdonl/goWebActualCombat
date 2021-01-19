package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int
var rw sync.RWMutex

func main() {
	ch := make(chan struct{}, 6)
	for i := 0; i < 3; i++ {
		go ReadCount(i, ch)
	}
	for i := 0; i < 3; i++ {
		go WriteCount(i, ch)
	}
	for i := 0; i < 6; i++ {
		<-ch
	}
}
func ReadCount(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}
func WriteCount(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(10)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}
