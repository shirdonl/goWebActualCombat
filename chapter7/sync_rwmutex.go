package main

import (
	"fmt"
	"sync"
)

var sum int
var mutex sync.RWMutex

func Reader() {
	mutex.RLock() //读锁
	fmt.Printf("读到的是%d\n", sum)
	defer mutex.RUnlock() //释放读锁
}

func Write() {
	mutex.Lock()
	defer mutex.Unlock()
	sum += 2
}

func main() {
	for x := 0; x < 10; x++ {
		go Reader()
	}
	for x := 0; x < 10; x++ {
		go Write()
	}
	for x := 0; x < 10; x++ {
		go Write()
	}
}