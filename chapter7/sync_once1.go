package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("test only once，这里只打印一次！")//打印
	}
	done := make(chan bool)
	for i := 0; i < 6; i++ {
		go func() {
			once.Do(onceBody)//确保只背执行一次
			done <- true
		}()
	}
	for i := 0; i < 6; i++ {
		<-done
	}
}
