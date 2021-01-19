package main

import (
	"fmt"
	"time"
)

func Timer(duration time.Duration) chan bool {
	ch := make(chan bool)

	go func() {
		time.Sleep(duration)
		// 到时间啦
		ch <- true
	}()

	return ch

}

func main() {
	// 定时5秒
	timeout := Timer(5 * time.Second)

	for {
		select {
		case <-timeout:
			// 到时
			fmt.Println("already 5s!")
			//结束程序
			return
		}
	}
}
